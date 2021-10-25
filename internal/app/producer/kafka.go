package producer

import (
	"github.com/ozonmp/bss-office-api/internal/app/repo"
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration // неиспользуется

	repo repo.EventRepo

	sender sender.EventSender
	events <-chan model.OfficeEvent

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	done chan bool
}

func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	repo repo.EventRepo,
	events <-chan model.OfficeEvent,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		repo:       repo,
		workerPool: workerPool,
		wg:         wg,
		done:       done,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					err := p.sender.Send(&event)
					if err != nil {
						p.workerPool.Submit(func() {
							err = p.repo.Unlock([]uint64{event.ID})
							if err != nil {
								log.Printf("produser unlock error:%s \n", err)
								// и чё делать?
							}
						})

						break
					}

					p.workerPool.Submit(func() {
						err = p.repo.Remove([]uint64{event.ID})
						if err != nil {
							log.Printf("produser remove error:%s \n", err)
							// и чё делать?
						}
					})
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
