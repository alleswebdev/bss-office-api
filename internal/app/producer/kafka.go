package producer

import (
	"context"
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
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	repo repo.EventRepo

	sender sender.EventSender
	events <-chan model.OfficeEvent

	workerPool *workerpool.WorkerPool

	wg *sync.WaitGroup
}

func NewKafkaProducer(
	n uint64,
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
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
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

func (p *producer) processClean(eventIDs []uint64) {
	err := p.repo.Remove(eventIDs)
	if err != nil {
		log.Printf("produser remove error:%s \n", err)
	}
}
