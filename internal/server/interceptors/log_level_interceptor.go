package interceptors

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

// ChangeLogLevelInterceptor позволяет изменить уровень логирование через мета-тег log-level
func ChangeLogLevelInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		mData, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return handler(ctx, req)
		}

		levels := mData.Get("log-level")

		if len(levels) == 0 {
			return handler(ctx, req)
		}

		level, ok := parseLevel(levels[0])
		if ok {
			newLogger := logger.CloneWithLevel(ctx, level)
			ctx = logger.AttachLogger(ctx, newLogger)
			logger.DebugKV(ctx, "ChangeLogLevelInterceptor - log level has been changed", "level", level)
		}

		return handler(ctx, req)
	}
}

func parseLevel(level string) (zapcore.Level, bool) {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel, true
	case "warning":
		return zap.WarnLevel, true
	case "error":
		return zap.ErrorLevel, true
	default:
		return zap.InfoLevel, false
	}

}
