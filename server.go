package main

import (
	"flag"
	"fmt"
	"github.com/PulseyTeam/game-server/handler"
	pb "github.com/PulseyTeam/game-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 3000, "The port to listen on.")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMultiplayerServiceServer(grpcServer, handler.NewMultiplayer())

	log.Printf("starting server on port %d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
