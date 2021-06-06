package main

import (
	"context"
	pb "github.com/PulseyTeam/game-server/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewMultiplayerServiceClient(conn)

	response, err := client.RoomConnect(ctx, &pb.RoomConnectRequest{MapId: "mehmetin_mapi"})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	log.Printf("Greeting: %s", response.GetRoomId())
}
