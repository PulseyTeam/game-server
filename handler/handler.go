package handler

import (
	"github.com/PulseyTeam/game-server/config"
	pb "github.com/PulseyTeam/game-server/proto"
	"github.com/PulseyTeam/game-server/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type MultiplayerHandler struct {
	pb.UnimplementedMultiplayerServiceServer
	rooms      map[string]map[string]*pb.Player
	jwtManager *service.JWTManager
	mongoDB    *mongo.Client
	cfg        *config.Config
}

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	jwtManager *service.JWTManager
	mongoDB    *mongo.Client
	cfg        *config.Config
}

func NewMultiplayerHandler(cfg *config.Config, mongoDB *mongo.Client, jwtManager *service.JWTManager) *MultiplayerHandler {
	return &MultiplayerHandler{cfg: cfg, mongoDB: mongoDB, jwtManager: jwtManager}
}

func NewAuthHandler(cfg *config.Config, mongoDB *mongo.Client, jwtManager *service.JWTManager) *AuthHandler {
	return &AuthHandler{cfg: cfg, mongoDB: mongoDB, jwtManager: jwtManager}
}
