package handler

import (
	"context"
	"github.com/PulseyTeam/game-server/model"
	pb "github.com/PulseyTeam/game-server/proto"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AuthHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	collection := a.mongoDB.Database(a.cfg.MongoDB.DB).Collection(model.UserCollection)

	findUser := &model.User{}
	filter := bson.D{primitive.E{Key: "username", Value: request.GetUsername()}}

	err := collection.FindOne(ctx, filter).Decode(&findUser)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed")
	}

	passwordCheck := findUser.CheckPassword(request.GetPassword())

	if !passwordCheck {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed")
	}

	accessToken, err := a.jwtManager.Generate(findUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &pb.LoginResponse{AccessToken: accessToken}, nil
}

func (a *AuthHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	collection := a.mongoDB.Database(a.cfg.MongoDB.DB).Collection(model.UserCollection)

	findUser := model.User{}
	filter := bson.D{primitive.E{Key: "username", Value: request.GetUsername()}}

	err := collection.FindOne(ctx, filter).Decode(&findUser)
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "user already registered")
	} else if err != mongo.ErrNoDocuments {
		log.Fatal().Err(err).Send()
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	newUser := &model.User{
		ID:       primitive.NewObjectID(),
		Username: request.GetUsername(),
		Password: request.GetPassword(),
	}

	err = newUser.HashPassword()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not hash password")
	}

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create resource")
	}

	accessToken, err := a.jwtManager.Generate(newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &pb.RegisterResponse{AccessToken: accessToken}, nil
}
