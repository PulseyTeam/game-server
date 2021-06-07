package main

import (
	"context"
	"fmt"
	"github.com/PulseyTeam/game-server/config"
	"github.com/PulseyTeam/game-server/database"
	"github.com/PulseyTeam/game-server/handler"
	"github.com/PulseyTeam/game-server/jwt"
	pb "github.com/PulseyTeam/game-server/proto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

func main() {
	if os.Getenv("KUBERNETES_PORT") != "" {
		zerolog.LevelFieldName = "severity"
		zerolog.TimestampFieldName = "timestamp"
		zerolog.TimeFieldFormat = time.RFC3339Nano
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}
		consoleWriter.TimeFormat = time.RFC3339
		log.Logger = zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()
	}

	log.Info().Msg("starting server...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	zerolog.SetGlobalLevel(zerolog.Level(cfg.Server.LogLevel))

	mongoDBConn, err := database.NewMongoDBConn(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer func() {
		if err := mongoDBConn.Disconnect(ctx); err != nil {
			log.Fatal().Err(err).Send()
		}
	}()
	log.Info().Msgf("mongodb connected: %v", mongoDBConn.NumberSessionsInProgress())

	jwtManager := jwt.NewManager()

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamInterceptor))
	pb.RegisterMultiplayerServiceServer(grpcServer, handler.NewMultiplayerHandler(cfg, mongoDBConn, jwtManager))
	pb.RegisterAuthServiceServer(grpcServer, handler.NewAuthHandler(cfg, mongoDBConn, jwtManager))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.Port))
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	log.Info().Msgf("server started on port %v", cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	now := time.Now()

	defer func() {
		log.Trace().Str("requestUrl", info.FullMethod).Str("latency", fmt.Sprintf("%v", time.Since(now))).Msgf("request completed")
	}()

	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	now := time.Now()

	defer func() {
		log.Trace().Str("requestUrl", info.FullMethod).Str("latency", fmt.Sprintf("%v", time.Since(now))).Msgf("stream completed")
	}()

	return handler(srv, stream)
}
