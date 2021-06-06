package handler

import (
	"github.com/PulseyTeam/game-server/config"
	pb "github.com/PulseyTeam/game-server/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type MultiplayerHandler struct {
	pb.UnimplementedMultiplayerServiceServer
	rooms   map[string]map[string]*pb.Player
	mongoDB *mongo.Client
	cfg     *config.Config
}

func NewMultiplayer(cfg *config.Config, mongoDB *mongo.Client) *MultiplayerHandler {
	return &MultiplayerHandler{cfg: cfg, mongoDB: mongoDB}
}
