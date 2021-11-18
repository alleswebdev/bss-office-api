package logger

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/bss-office-api/internal/config"
	gelf "github.com/snovichkov/zap-gelf"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap/zapcore"
	"log"
	"os"

	"go.uber.org/zap"
)

type ctxKey struct{}

var attachedLoggerKey = &ctxKey{}

var globalLogger *zap.SugaredLogger

func fromContext(ctx context.Context) *zap.SugaredLogger {
	result := globalLogger

	if attachedLogger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		result = attachedLogger
	}

	// добавим id спана в лог
	jaegerSpan := opentracing.SpanFromContext(ctx)
	if jaegerSpan != nil {
		if spanCtx, ok := opentracing.SpanFromContext(ctx).Context().(jaeger.SpanContext); ok {
			result = result.With("trace-id", spanCtx.TraceID())
		}
	}

	return result
}

// ErrorKV - создаёт лог с уровнем error
func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Errorw(message, kvs...)
}

// WarnKV - создаёт лог с уровнем warning
func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Warnw(message, kvs...)
}

// InfoKV - создаёт лог с уровнем info
func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Infow(message, kvs...)
}

// DebugKV - создаёт лог с уровнем debug
func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Debugw(message, kvs...)
}

// FatalKV - создаёт лог с уровнем fatal
func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Fatalw(message, kvs...)
}

// AttachLogger - добавляет логгер к контексту
func AttachLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, attachedLoggerKey, logger)
}

// CloneWithLevel - клонирует существующий логгер с указанным уровнем логирования
func CloneWithLevel(ctx context.Context, newLevel zapcore.Level) *zap.SugaredLogger {
	return fromContext(ctx).
		Desugar().
		WithOptions(WithLevel(newLevel)).
		Sugar()
}

// SetLogger - устанавливает глобальный логгер
func SetLogger(newLogger *zap.SugaredLogger) {
	globalLogger = newLogger
}

func init() {
	notSugaredLogger, err := zap.NewProduction()
	if err != nil {
		log.Panic(err)
	}

	globalLogger = notSugaredLogger.Sugar()
}

func InitLogger(ctx context.Context, cfg *config.Config) (syncFn func()) {
	loggingLevel := zap.InfoLevel

	if cfg.Project.Debug {
		loggingLevel = zap.DebugLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(cfg.Telemetry.GraylogPath),
		gelf.Level(loggingLevel),
	)

	if err != nil {
		FatalKV(ctx, "logger create error", "err", err)
	}

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

	sugaredLogger := notSugaredLogger.Sugar()
	SetLogger(sugaredLogger.With(
		"service", cfg.Project.Name,
	))

	return func() {
		errInit := notSugaredLogger.Sync()
		if errInit != nil {
			ErrorKV(ctx, "initLogger() error", "err", errInit)
		}
	}
}
