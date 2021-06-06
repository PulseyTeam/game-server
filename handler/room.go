package handler

import (
	"context"
	"github.com/PulseyTeam/game-server/model"
	pb "github.com/PulseyTeam/game-server/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func (h *MultiplayerHandler) RoomConnect(ctx context.Context, request *pb.RoomConnectRequest) (*pb.RoomConnectResponse, error) {
	collection := h.mongoDB.Database(h.cfg.MongoDB.DB).Collection("game_sessions")

	gameSession := model.GameSession{}
	filter := bson.D{
		primitive.E{Key: "map_id", Value: request.GetMapId()},
		primitive.E{Key: "status", Value: model.SessionWaiting},
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&gameSession)
	if err == nil {
		return &pb.RoomConnectResponse{RoomId: gameSession.ID.String()}, nil
	} else {
		log.Printf("find error: %v", err)
	}

	result, err := collection.InsertOne(context.TODO(), model.GameSession{
		ID:         primitive.NewObjectID(),
		MapID:      request.GetMapId(),
		Status:     model.SessionWaiting,
		StartedAt:  time.Now(),
		FinishedAt: nil,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create resource")
	}

	insertedID := result.InsertedID.(primitive.ObjectID).String()

	log.Printf("game session (created): %v", insertedID)

	return &pb.RoomConnectResponse{RoomId: insertedID}, nil
}

func (h *MultiplayerHandler) RoomStream(stream pb.MultiplayerService_RoomStreamServer) error {
	if h.rooms == nil {
		h.rooms = make(map[string]*pb.RoomStreamRequest)
	}

	for {
		request, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		h.rooms[request.GetRoomId()] = &pb.RoomStreamRequest{
			Player: request.GetPlayer(),
			RoomId: request.GetRoomId(),
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
