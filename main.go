package main

import (
	"context"
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

func (s *server) SetPosition(ctx context.Context, request *pb.SetPositionRequest) (*pb.SetPositionResponse, error) {
	if s.players == nil {
		s.players = make(map[string]*pb.PlayerPosition)
	}

	s.players[request.GetName()] = &pb.PlayerPosition{
		Name: request.GetName(),
		X:    request.GetX(),
		Y:    request.GetY(),
	}

	log.Printf("SetPosition: %v -> %v, %v", request.GetName(), request.GetX(), request.GetY())

	return &pb.SetPositionResponse{Success: true}, nil
}

func (s *server) GetPositions(ctx context.Context, request *pb.GetPositionsRequest) (*pb.GetPlayerPositions, error) {
	currentPlayers := make([]*pb.PlayerPosition, 0, len(s.players))

	for _, player := range s.players {
		currentPlayers = append(currentPlayers, player)
	}

	return &pb.GetPlayerPositions{Players: currentPlayers}, nil
}

func (s *server) BiDirectSetPositions(stream pb.MultiplayerService_BiDirectSetPositionsServer) error {
	if s.players == nil {
		s.players = make(map[string]*pb.PlayerPosition)
	}

	for {
		request, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		s.players[request.GetName()] = &pb.PlayerPosition{
			Name: request.GetName(),
			X:    request.GetX(),
			Y:    request.GetY(),
		}
		log.Printf("SetPosition: %v -> %v, %v", request.GetName(), request.GetX(), request.GetY())

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
