package handler

import (
	"github.com/PulseyTeam/game-server/config"
	"github.com/PulseyTeam/game-server/jwt"
	pb "github.com/PulseyTeam/game-server/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type MultiplayerHandler struct {
	pb.UnimplementedMultiplayerServiceServer
	rooms      map[string]map[string]*pb.Player
	jwtManager *jwt.Manager
	mongoDB    *mongo.Client
	cfg        *config.Config
}

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	jwtManager *jwt.Manager
	mongoDB    *mongo.Client
	cfg        *config.Config
}

func NewMultiplayerHandler(cfg *config.Config, mongoDB *mongo.Client, jwtManager *jwt.Manager) *MultiplayerHandler {
	return &MultiplayerHandler{cfg: cfg, mongoDB: mongoDB, jwtManager: jwtManager}
}

func NewAuthHandler(cfg *config.Config, mongoDB *mongo.Client, jwtManager *jwt.Manager) *AuthHandler {
	return &AuthHandler{cfg: cfg, mongoDB: mongoDB, jwtManager: jwtManager}
}
