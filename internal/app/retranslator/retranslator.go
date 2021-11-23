// Package retranslator  предназначен для пересылки событий из репозитория в брокер сообщений
package retranslator

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/app/consumer"
	"github.com/ozonmp/bss-office-api/internal/app/producer"
	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/config"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/ozonmp/bss-office-api/internal/repo"

	"github.com/gammazero/workerpool"
)

// Retranslator interface
type Retranslator interface {
	Start(ctx context.Context)
	Close()
}

type retranslator struct {
	events     chan model.OfficeEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

// NewRetranslator create new Retranslator from config
func NewRetranslator(cfg *config.Config, repo repo.EventRepo, sender sender.EventSender) Retranslator {
	events := make(chan model.OfficeEvent, cfg.Retranslator.ChannelSize)
	workerPool := workerpool.New(cfg.Retranslator.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.Retranslator.ConsumerCount,
		cfg.Retranslator.ConsumeSize,
		cfg.Retranslator.ConsumeTimeout,
		repo,
		events)
	producer := producer.NewKafkaProducer(
		cfg.Retranslator.ProducerCount,
		cfg.Retranslator.ProducerBatchSize,
		cfg.Retranslator.ProducerTimeout,
		sender,
		repo,
		events,
		workerPool)

	return &retranslator{
		events:     events,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

// Start запускает ранее сконфигурированный ретранслятор и его продюсееров и консюмеров
func (r *retranslator) Start(ctx context.Context) {
	r.producer.StartBatch(ctx)
	r.consumer.Start(ctx)
}

// Close останавливает работу ретрянслятора, его продюсеров, консюмеров и воркерпул
func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
