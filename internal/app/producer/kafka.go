// Package producer предназначен для пересылки событий в брокер сообщений
package producer

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/ozonmp/bss-office-api/internal/repo"
	"log"
	"sync"
	"time"

	"github.com/gammazero/workerpool"
)

// Producer interface
type Producer interface {
	Start(ctx context.Context)
	StartBatch(ctx context.Context)
	Close()
}

var errUnlock = errors.New("producer: unlock error: %s")
var errRemove = errors.New("producer: remove error: %s")
var errBatchHandler = errors.New("producer: batch handler error: %s")
var errChannelClose = errors.New("producer: consumer closed the channel ")
var errSender = errors.New("producer: error send event: %s")

type producer struct {
	n int

	repo      repo.EventRepo
	batchSize int

	timeout time.Duration
	sender sender.EventSender
	events <-chan model.OfficeEvent

	workerPool *workerpool.WorkerPool

	wg *sync.WaitGroup
}

// NewKafkaProducer create a new kafka producer
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
						fmt.Println(errChannelClose)
						return
					}

					err := p.sender.Send(ctx, &event)
					if err != nil {
						p.processUpdate(ctx, []uint64{event.ID})
						continue
					}

					p.processClean(ctx, []uint64{event.ID})
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

			updateChannel := p.startBatchHandler(ctx, p.processWaitUpdate)
			removeChannel := p.startBatchHandler(ctx, p.processWaitClean)
			defer close(updateChannel)
			defer close(removeChannel)

			for {
				select {
				case event, ok := <-p.events:
					if !ok {
						log.Println(errChannelClose)
						return
					}

					err := p.sender.Send(ctx, &event)
					if err != nil {
						log.Printf(errSender.Error(), err)
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

func (p *producer) processUpdate(ctx context.Context, eventIDs []uint64) {
	p.workerPool.Submit(func() {
		err := p.repo.Unlock(ctx, eventIDs)
		if err != nil {
			log.Printf(errUnlock.Error(), err)
		}
	})
}

// processWaitUpdate разблокирует записи в репозитории и дожидается возврата ошибки
func (p *producer) processWaitUpdate(ctx context.Context, eventIDs []uint64) error {
	errChan := make(chan error)
	defer close(errChan)
	p.workerPool.Submit(func() {
		err := p.repo.Unlock(ctx, eventIDs)
		if err != nil {
			errChan <- fmt.Errorf(errUnlock.Error(), err)
			return
		}

		errChan <- nil
	})

	return <-errChan
}

func (p *producer) processClean(ctx context.Context, eventIDs []uint64) {
	err := p.repo.Remove(ctx, eventIDs)
	if err != nil {
		log.Printf(errRemove.Error(), err)
	}
}

// processWaitClean удаляет обработанные записи в репозитории и дожидается возврата ошибки
func (p *producer) processWaitClean(ctx context.Context, eventIDs []uint64) error {
	errChan := make(chan error)
	defer close(errChan)

	p.workerPool.Submit(func() {
		err := p.repo.Remove(ctx, eventIDs)
		if err != nil {
			errChan <- fmt.Errorf(errRemove.Error(), err)
			return
		}

		errChan <- nil
	})

	return <-errChan
}

// startBatchHandler предназначен для пакетного удаления и обновления событий в репозитории
// вторым аргументом передаётся функция для удаления/обновления событий (processWaitUpdate или processWaitClean)
// startBatchHandler слушает канал для событий, id которых необходимо обработать
// и складывает их в буфер размером Producer.batchSize, при наполнении буфера,
// завершении контекста, или по таймауту Producer.timeout
// буфер отсылается в пул воркеров для обработки через переданную функцию
// Для остановки необходимо завершить переданый контекст
//
// startBatchHandler возвращает канал для работы с событиями, вызывающий функцию должен закрыть этот канал
// при завершении работы
func (p *producer) startBatchHandler(ctx context.Context, f func(ctx context.Context, ids []uint64) error) chan<- uint64 {
	c := make(chan uint64)

	buffer := make([]uint64, 0, p.batchSize)
	ticker := time.NewTicker(p.timeout)

	go func() {
		for {
			select {
			case id, ok := <-c:
				if !ok {
					if len(buffer) > 0 {
						err := f(ctx, buffer)
						if err != nil {
							log.Printf(errBatchHandler.Error(), err)
						}
					}

					ticker.Stop()
					log.Println("batch handler channel was closed")
					return
				}

				buffer = append(buffer, id)

				if len(buffer) >= p.batchSize {
					err := f(ctx, buffer)

					if err != nil {
						log.Printf(errBatchHandler.Error(), err)
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
				err := f(ctx, buffer)

				if err != nil {
					log.Printf(errBatchHandler.Error(), err)
					continue
				}

				buffer = buffer[:0]
			case <-ctx.Done():
				ticker.Stop()
				if len(buffer) != 0 {
					err := f(ctx, buffer)
					if err != nil {
						log.Printf(errBatchHandler.Error(), err)
					}
				}

				return
			}
		}
	}()

	return c
}
