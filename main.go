package main

import (
	"flag"
	"fmt"
	pb "github.com/PulseyTeam/game-server/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedMultiplayerServiceServer
	players map[string]*pb.PlayerPosition
}

func (s *server) BiDirectSetPositions(stream pb.MultiplayerService_BiDirectSetPositionsServer) error {
	if s.players == nil {
		s.players = make(map[string]*pb.PlayerPosition)
	}

	var tickCounter uint64 = 0

	for {
		request, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		tickCounter++

		s.players[request.GetName()] = &pb.PlayerPosition{
			Name: request.GetName(),
			X:    request.GetX(),
			Y:    request.GetY(),
		}

		if (tickCounter % 20) == 0 {
			log.Printf("New Position: %v (%v, %v)", request.GetName(), request.GetX(), request.GetY())
		}

		currentPlayers := make([]*pb.PlayerPosition, 0, len(s.players))
		for _, player := range s.players {
			currentPlayers = append(currentPlayers, player)
		}

		err = stream.Send(&pb.GetPlayerPositions{Players: currentPlayers})
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	return nil
}

func main() {
	port := flag.Int("port", 3000, "The port to listen on.")
	flag.Parse()

	log.Printf("listening on port %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMultiplayerServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
