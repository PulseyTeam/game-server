package interceptor

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"time"
)

func Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()

	defer func() {
		httpRequest := httpRequest{
			RequestUrl: info.FullMethod,
			Latency:    fmt.Sprintf("%fs", time.Since(start).Seconds()),
		}
		log.Trace().Interface("httpRequest", httpRequest).Msgf("request completed")
	}()

	return handler(ctx, req)
}
