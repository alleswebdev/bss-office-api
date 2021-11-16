package interceptors

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

// VerboseRequestInterceptor - логгирует информацию о запросе и ответе при получение мета тега verbose=true
func VerboseRequestInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		mData, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return handler(ctx, req)
		}

		verboseTag := mData.Get("verbose")
		isVerbose, err := strconv.ParseBool(verboseTag[0])

		if !isVerbose || err != nil {
			return handler(ctx, req)
		}

		logger.DebugKV(ctx, "Request debug info",
			"RPC method", info.FullMethod,
			"metadata", mData,
			"request", req,
		)

		response, err := handler(ctx, req)

		if err == nil {
			logger.DebugKV(ctx, "Response debug info",
				"RPC method", info.FullMethod,
				"metadata", mData,
				"Response", response,
			)
		}

		return response, err
	}
}
