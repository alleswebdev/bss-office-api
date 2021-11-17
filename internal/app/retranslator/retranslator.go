// Package retranslator  предназначен для пересылки событий из репозитория в брокер сообщений
package retranslator

import (
	"context"
	"time"

	"github.com/ozonmp/bss-office-api/internal/app/consumer"
	"github.com/ozonmp/bss-office-api/internal/app/producer"
	"github.com/ozonmp/bss-office-api/internal/app/repo"
	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/model"

	"github.com/gammazero/workerpool"
)

// Retranslator interface
type Retranslator interface {
	Start(ctx context.Context)
	Close()
}

// Config конфигурирует запускаемый экземпляр ретранслятора
// ChannelSize - размер канала пересылаемых событий
// ConsumerCount - количество горутин с Consumer
// ConsumeSize  - размер канала Consumer
// ConsumeTimeout - время ожидания до следующего batch-запроса
// ProducerCount - количество горутин с Producer
// WorkerCount - количество воркеров в воркерпуле для очистки и обновления событий в БД после отправки в брокер сообщений
// Repo - репозиторий для работы с событиями
// Sender - сервис для отправки событий в кафку
type Config struct {
	ChannelSize uint64

	ConsumerCount  int
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount     int
	ProducerTimeout   time.Duration
	ProducerBatchSize int
	WorkerCount       int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan model.OfficeEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

// NewRetranslator create new Retranslator from config
func NewRetranslator(cfg Config) Retranslator {
	events := make(chan model.OfficeEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events)
	producer := producer.NewKafkaProducer(
		cfg.ProducerCount,
		cfg.ProducerBatchSize,
		cfg.ProducerTimeout,
		cfg.Sender,
		cfg.Repo,
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
	r.producer.Start(ctx)
	r.consumer.Start(ctx)
}

// Close останавливает работу ретрянслятора, его продюсеров, консюмеров и воркерпул
func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
