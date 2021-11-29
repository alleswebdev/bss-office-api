package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozonmp/bss-office-api/internal/app/metrics"
	"github.com/ozonmp/bss-office-api/internal/app/retranslator"
	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/config"
	"github.com/ozonmp/bss-office-api/internal/database"
	"github.com/ozonmp/bss-office-api/internal/logger"
	"github.com/ozonmp/bss-office-api/internal/repo"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	sigs := make(chan os.Signal, 1)

	if err := config.ReadConfigYML("config.retranslator.yml"); err != nil {
		log.Fatalf("Failed init configuration:%s", err)
	}

	cfg := config.GetConfigInstance()

	syncLogger := logger.InitLogger(ctx, &cfg)
	defer syncLogger()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(ctx, dsn, cfg.Database.Driver)
	if err != nil {
		logger.FatalKV(ctx, "Failed init postgres", "err", err)
	}
	defer db.Close()

	eventsRepo := repo.NewEventRepo(db)
	sender, err := sender.NewEventSender(cfg.Kafka.Brokers, cfg.Kafka.Topic)

	if err != nil {
		logger.FatalKV(ctx, "Failed start kafka producer", "err", err)
	}

	retranslator := retranslator.NewRetranslator(&cfg, eventsRepo, sender)
	defer retranslator.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	metricsServer := metrics.CreateMetricsServer(&cfg)

	go func(ctx context.Context) {
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.FatalKV(ctx, "Failed running metrics server", "err", err)
			cancel()
		}
	}(ctx)

	metrics.InitMetrics(cfg)
	retranslator.Start(ctx)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
