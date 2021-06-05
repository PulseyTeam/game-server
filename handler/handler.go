package handler

import (
	pb "github.com/PulseyTeam/game-server/proto"
)

type MultiplayerHandler struct {
	pb.UnimplementedMultiplayerServiceServer
	rooms map[int64]*pb.Player
}

func NewMultiplayer() *MultiplayerHandler {
	return &MultiplayerHandler{}
}
