package handler

import (
	"context"
	"github.com/PulseyTeam/game-server/model"
	pb "github.com/PulseyTeam/game-server/proto"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"time"
)

func (h *MultiplayerHandler) RoomConnect(ctx context.Context, request *pb.RoomConnectRequest) (*pb.RoomConnectResponse, error) {
	collection := h.mongoDB.Database(h.cfg.MongoDB.DB).Collection(model.GameSessionCollection)

	gameSession := model.GameSession{}
	filter := bson.D{
		primitive.E{Key: "map_id", Value: request.GetMapId()},
		primitive.E{Key: "status", Value: model.StatusWaiting},
	}

	err := collection.FindOne(ctx, filter).Decode(&gameSession)
	if err == nil {
		return &pb.RoomConnectResponse{RoomId: gameSession.ID.Hex()}, nil
	} else {
		//Todo refactor
		log.Warn().Msgf("find error: %v", err)
	}

	result, err := collection.InsertOne(ctx, model.GameSession{
		ID:         primitive.NewObjectID(),
		MapID:      request.GetMapId(),
		Status:     model.StatusWaiting,
		StartedAt:  time.Now(),
		FinishedAt: nil,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create resource")
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	log.Trace().Msgf("game session created: %v", insertedID)

	return &pb.RoomConnectResponse{RoomId: insertedID}, nil
}

func (h *MultiplayerHandler) RoomStream(stream pb.MultiplayerService_RoomStreamServer) error {
	if h.rooms == nil {
		h.rooms = make(map[string]map[string]*pb.Player)
	}

	for {
		request, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if h.rooms[request.GetRoomId()] == nil {
			h.rooms[request.GetRoomId()] = make(map[string]*pb.Player)
		}

		h.rooms[request.GetRoomId()][request.GetPlayer().GetId()] = request.GetPlayer()

		currentPlayers := make([]*pb.Player, 0, len(h.rooms[request.GetRoomId()]))
		for _, player := range h.rooms[request.GetRoomId()] {
			currentPlayers = append(currentPlayers, player)
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
