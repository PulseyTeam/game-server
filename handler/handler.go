package handler

import (
	pb "github.com/PulseyTeam/game-server/proto"
)

type MultiplayerHandler struct {
	pb.UnimplementedMultiplayerServiceServer
	rooms map[uint64]*pb.RoomStreamRequest
}

func NewMultiplayer() *MultiplayerHandler {
	return &MultiplayerHandler{}
}
