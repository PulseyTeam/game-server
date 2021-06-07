package interceptor

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"time"
)

func Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()

	defer func() {
		httpRequest := httpRequest{
			RequestUrl: info.FullMethod,
			Latency:    fmt.Sprintf("%fs", time.Since(start).Seconds()),
		}
		log.Trace().Interface("httpRequest", httpRequest).Msgf("stream completed")
	}()

	return handler(srv, stream)
}
