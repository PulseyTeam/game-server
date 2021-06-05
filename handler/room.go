package handler

import (
	"context"
	pb "github.com/PulseyTeam/game-server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

func (h *MultiplayerHandler) RoomConnect(ctx context.Context, request *pb.RoomConnectRequest) (*pb.RoomConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomConnect not implemented")
}

func (h *MultiplayerHandler) RoomStream(stream pb.MultiplayerService_RoomStreamServer) error {
	if h.rooms == nil {
		h.rooms = make(map[uint64]*pb.RoomStreamRequest)
	}

	for {
		request, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		h.rooms[request.RoomId] = &pb.RoomStreamRequest{
			Player: request.GetPlayer(),
			RoomId: request.RoomId,
		}

		currentPlayers := make([]*pb.Player, 0, len(h.rooms))
		for _, player := range h.rooms {
			if player.GetRoomId() != request.GetRoomId() {
				continue
			}
			currentPlayers = append(currentPlayers, player.GetPlayer())
		}

		err = stream.Send(&pb.RoomStreamResponse{Players: currentPlayers})
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	return nil
}
