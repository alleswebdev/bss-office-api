package main

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-office-api/internal/app/sender"
	"github.com/ozonmp/bss-office-api/internal/config"
	"github.com/ozonmp/bss-office-api/internal/database"
	"github.com/ozonmp/bss-office-api/internal/repo"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozonmp/bss-office-api/internal/app/retranslator"
)

func main() {

	sigs := make(chan os.Signal, 1)

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	cfg := config.GetConfigInstance()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close()

	rCfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  10,
		ConsumeSize:    10,
		ConsumeTimeout: time.Millisecond,
		ProducerCount:  28,
		WorkerCount:    10,
		Repo:           repo.NewEventRepo(db),
		Sender:         sender.NewDummySender(),
	}

	retranslator := retranslator.NewRetranslator(rCfg)
	defer retranslator.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	retranslator.Start(ctx)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
