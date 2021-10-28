package producer

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozonmp/bss-office-api/internal/app/repo"
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start(ctx context.Context)
	StartBatch(ctx context.Context)
	Close()
}

var BatchUnlockErr = errors.New("batch updater: unlock error: %s")

type producer struct {
	n       int
	timeout time.Duration

	repo      repo.EventRepo
	batchSize int

	sender sender.EventSender
	events <-chan model.OfficeEvent

	workerPool *workerpool.WorkerPool

	wg *sync.WaitGroup
}

func NewKafkaProducer(
	n int,
	batchSize int,
	timeout time.Duration,
	sender sender.EventSender,
	repo repo.EventRepo,
	events <-chan model.OfficeEvent,
	workerPool *workerpool.WorkerPool,
) Producer {

	var wg sync.WaitGroup

	return &producer{
		n:          n,
		batchSize:  batchSize,
		timeout:    timeout,
		sender:     sender,
		events:     events,
		repo:       repo,
		workerPool: workerPool,
		wg:         &wg,
	}
}

// Start Запускает Producer в single режиме,
// update и clean воркеры будут отправлять события в репо
// по одному сразу после получения
func (p *producer) Start(ctx context.Context) {
	for i := 0; i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()

			for {
				select {
				case event, ok := <-p.events:
					if !ok {
						fmt.Println("producer: consumer channel close")
						return
					}

					err := p.sender.Send(ctx, &event)
					if err != nil {
						p.processUpdate([]uint64{event.ID})
						continue
					}

					p.processClean([]uint64{event.ID})
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

// StartBatch Запускает Producer в batch режиме,
// update и clean воркеры будут накапливать id событий и отправлять
// в репо пачками, а не по одному
func (p *producer) StartBatch(ctx context.Context) {
	for i := 0; i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()

			updateChannel := p.startBatchUpdater(ctx)
			removeChannel := p.startBatchCleaner(ctx)

			for {
				select {
				case event, ok := <-p.events:
					if !ok {
						fmt.Println("producer: consumer channel close")
						return
					}

					err := p.sender.Send(ctx, &event)
					if err != nil {
						updateChannel <- event.ID
						continue
					}

					removeChannel <- event.ID
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	p.wg.Wait()
}

func (p *producer) processUpdate(eventIDs []uint64) {
	p.workerPool.Submit(func() {
		err := p.repo.Unlock(eventIDs)
		if err != nil {
			log.Printf("produser unlock error:%s \n", err)
		}
	})
}

// processWaitUpdate разблокирует записи в репозитории и дожидается возврата ошибки
func (p *producer) processWaitUpdate(eventIDs []uint64) error {
	errChan := make(chan error)
	p.workerPool.Submit(func() {
		err := p.repo.Unlock(eventIDs)
		if err != nil {
			errChan <- fmt.Errorf("produser unlock error:%s \n", err)
		}

		errChan <- nil
	})

	return <-errChan
}

// BatchUpdater слушает канал для  событий, id  которые необходимо разлочить
// и складывает их в буфер размером Producer.batchSize, при наполнении буфера,
// завершении контекста, или по итечению таймаута Producer.timeout
// буфер отсылается в пул воркеров для разблокирования
// Для остановки необходимо завершить переданый контекст
func (p *producer) startBatchUpdater(ctx context.Context) chan<- uint64 {
	c := make(chan uint64)

	buffer := make([]uint64, 0, p.batchSize)
	ticker := time.NewTicker(p.timeout)

	go func() {
		for {
			select {
			case id, ok := <-c:
				if !ok {
					err := p.processWaitUpdate(buffer)
					if err != nil {
						log.Printf(BatchUnlockErr.Error(), err)
					}
					ticker.Stop()
					log.Println("update channel was closed")
					return
				}

				buffer = append(buffer, id)

				if len(buffer) >= p.batchSize {
					err := p.processWaitUpdate(buffer)

					if err != nil {
						log.Printf(BatchUnlockErr.Error(), err)
						c <- id // вернём обратно, чтобы не потерять событие
						continue
					}

					buffer = buffer[:0]
					continue
				}
			case <-ticker.C:
				if len(buffer) == 0 {
					continue
				}
				err := p.processWaitUpdate(buffer)

				if err != nil {
					log.Printf(BatchUnlockErr.Error(), err)
					continue
				}

				buffer = buffer[:0]
			case <-ctx.Done():
				ticker.Stop()
				if len(buffer) != 0 {
					err := p.processWaitUpdate(buffer)
					if err != nil {
						log.Printf(BatchUnlockErr.Error(), err)
					}
				}

				return
			}
		}
	}()

	return c
}

func (p *producer) processClean(eventIDs []uint64) {
	err := p.repo.Remove(eventIDs)
	if err != nil {
		log.Printf("produser remove error:%s \n", err)
	}
}

// processWaitClean удаляет обработанные записи в репозитории и дожидается возврата ошибки
func (p *producer) processWaitClean(eventIDs []uint64) error {
	errChan := make(chan error)
	p.workerPool.Submit(func() {
		err := p.repo.Remove(eventIDs)
		if err != nil {
			errChan <- fmt.Errorf("produser unlock error:%s \n", err)
		}

		errChan <- nil
	})

	return <-errChan
}

// startBatchCleaner слушает канал для событий, id которых необходимо удалить
// и складывает их в буфер размером Producer.batchSize, при наполнении буфера,
// завершении контекста, или по итечению таймаута Producer.timeout
// буфер отсылается в пул воркеров для удаления
// Для остановки необходимо завершить переданый контекст
func (p *producer) startBatchCleaner(ctx context.Context) chan<- uint64 {
	c := make(chan uint64)

	buffer := make([]uint64, 0, p.batchSize)
	ticker := time.NewTicker(p.timeout)

	go func() {
		for {
			select {
			case id, ok := <-c:
				if !ok {
					err := p.processWaitClean(buffer)
					if err != nil {
						log.Printf(BatchUnlockErr.Error(), err)
					}
					ticker.Stop()
					log.Println("cleaner channel was closed")
					return
				}

				buffer = append(buffer, id)

				if len(buffer) >= p.batchSize {
					err := p.processWaitClean(buffer)

					if err != nil {
						log.Printf(BatchUnlockErr.Error(), err)
						c <- id // вернём обратно, чтобы не потерять событие
						continue
					}

					buffer = buffer[:0]
					continue
				}
			case <-ticker.C:
				if len(buffer) == 0 {
					continue
				}
				err := p.processWaitClean(buffer)

				if err != nil {
					log.Printf(BatchUnlockErr.Error(), err)
					continue
				}

				buffer = buffer[:0]
			case <-ctx.Done():
				ticker.Stop()
				if len(buffer) != 0 {
					err := p.processWaitClean(buffer)
					if err != nil {
						log.Printf(BatchUnlockErr.Error(), err)
					}
				}

				return
			}
		}
	}()

	return c
}
