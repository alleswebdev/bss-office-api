package producer

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/ozonmp/bss-office-api/internal/repo"
	"log"
	"sync"

	"github.com/gammazero/workerpool"
)

// Producer interface
type Producer interface {
	Start(ctx context.Context)
	Close()
}

type producer struct {
	n int

	repo repo.EventRepo

	sender sender.EventSender
	events <-chan model.OfficeEvent

	workerPool *workerpool.WorkerPool

	wg *sync.WaitGroup
}

// NewKafkaProducer create a new kafka producer
func NewKafkaProducer(
	n int,
	sender sender.EventSender,
	repo repo.EventRepo,
	events <-chan model.OfficeEvent,
	workerPool *workerpool.WorkerPool,
) Producer {

	var wg sync.WaitGroup

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		repo:       repo,
		workerPool: workerPool,
		wg:         &wg,
	}
}

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

func (p *producer) Close() {
	p.wg.Wait()
}

func (p *producer) processUpdate(ctx context.Context, eventIDs []uint64) {
	p.workerPool.Submit(func() {
		err := p.repo.Unlock(ctx, eventIDs)
		if err != nil {
			log.Printf("produser unlock error:%s \n", err)
		}
	})
}

func (p *producer) processClean(ctx context.Context, eventIDs []uint64) {
	err := p.repo.Remove(ctx, eventIDs)
	if err != nil {
		log.Printf("produser remove error:%s \n", err)
	}
}
