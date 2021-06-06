package main

import (
	"context"
	"fmt"
	"github.com/PulseyTeam/game-server/config"
	"github.com/PulseyTeam/game-server/database"
	"github.com/PulseyTeam/game-server/handler"
	pb "github.com/PulseyTeam/game-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("starting server...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	mongoDBConn, err := database.NewMongoDBConn(ctx, cfg)
	if err != nil {
		log.Fatalf("cannot connect mongodb: %v", err)
	}
	defer func() {
		if err := mongoDBConn.Disconnect(ctx); err != nil {
			log.Fatalf("mongodb disconnect: %v", err)
		}
	}()
	log.Printf("mongodb connected: %v", mongoDBConn.NumberSessionsInProgress())

	grpcServer := grpc.NewServer()
	pb.RegisterMultiplayerServiceServer(grpcServer, handler.NewMultiplayer(cfg, mongoDBConn))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server started on port %v", cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
