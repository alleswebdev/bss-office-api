// Package consumer  получает модели событий из репозитория и пересылает продюсерам
package consumer

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/repo"
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-office-api/internal/model"
)

// Consumer interface
type Consumer interface {
	Start(ctx context.Context)
	Close()
}

type consumer struct {
	n      int
	events chan<- model.OfficeEvent

	repo      repo.EventRepo
	batchSize uint64
	timeout   time.Duration

	wg *sync.WaitGroup
}

// NewDbConsumer create a new consumer
// n - количество гоурутин с консюмерами
// batchSize - количество получаемых событий из репозитория за раз
// consumeTimeout - периодичность получение событий из репозитория
// repo - репозиторий для работы
// events - канал для работы с продюсером, события будут записываться в него
//
// при завершении работы консюмер закроет канал events
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
					events, err := c.repo.Lock(ctx, c.batchSize)
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
