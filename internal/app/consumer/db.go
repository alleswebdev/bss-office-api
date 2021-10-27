package consumer

import (
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-office-api/internal/app/repo"
	"github.com/ozonmp/bss-office-api/internal/model"
)

type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	n      int
	events chan<- model.OfficeEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	done chan struct{}
	wg   *sync.WaitGroup
}

func NewDbConsumer(
	n int,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- model.OfficeEvent) Consumer {

	var wg sync.WaitGroup
	done := make(chan struct{})

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        &wg,
		done:      done,
	}
}

func (c *consumer) Start() {
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
				case <-c.done:
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	close(c.done)
	c.wg.Wait()
	close(c.events)
}
