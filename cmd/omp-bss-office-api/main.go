package main

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-office-api/internal/app/metrics"
	"github.com/ozonmp/bss-office-api/internal/app/retranslator"
	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/config"
	"github.com/ozonmp/bss-office-api/internal/database"
	"github.com/ozonmp/bss-office-api/internal/repo"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	sigs := make(chan os.Signal, 1)

	cfg := config.GetConfigInstance()

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
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close()

	eventsRepo := repo.NewEventRepo(db)
	dummySender := sender.NewDummySender()

	retranslator := retranslator.NewRetranslator(cfg, eventsRepo, dummySender)
	defer retranslator.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	metrics.InitMetrics(cfg)
	retranslator.Start(ctx)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
