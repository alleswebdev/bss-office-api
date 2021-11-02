package consumer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-office-api/internal/app/repo"
	"github.com/ozonmp/bss-office-api/internal/model"
)

type Consumer interface {
	Start(ctx context.Context)
	Close()
}

type consumer struct {
	n      int
	events chan<- model.OfficeEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	wg *sync.WaitGroup
}

func NewDbConsumer(
	n int,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- model.OfficeEvent) Consumer {

	var wg sync.WaitGroup

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        &wg,
	}
}

func (c *consumer) Start(ctx context.Context) {
	for i := 0; i < c.n; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(c.batchSize)
					if err != nil {
						log.Printf("consumer lock error:%s \n", err)
						continue
					}
					for _, event := range events {
						c.events <- event
					}
				case <-ctx.Done():
					ticker.Stop()
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	c.wg.Wait()
	close(c.events)
}
