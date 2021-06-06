package main

import (
	"context"
	"github.com/PulseyTeam/game-server/handler"
	pb "github.com/PulseyTeam/game-server/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

func newTestServer(ctx context.Context) (pb.MultiplayerServiceClient, func()) {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()
	pb.RegisterMultiplayerServiceServer(grpcServer, &handler.MultiplayerHandler{})
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, _ := grpc.DialContext(ctx, "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())

	closer := func() {
		lis.Close()
		grpcServer.Stop()
	}

	client := pb.NewMultiplayerServiceClient(conn)

	return client, closer
}

func TestRoomStream(t *testing.T) {
	type expectation struct {
		count int
		out   *pb.RoomStreamResponse
		err   error
	}

	tcs := map[string]struct {
		in       *pb.RoomStreamRequest
		expected expectation
	}{
		"ok": {
			in: &pb.RoomStreamRequest{
				Player: &pb.Player{Id: "1", Name: "RexiusTR", Position: &pb.Coordinate{X: 66, Y: 87}, Direction: pb.Direction_STOP},
				RoomId: "60bc207d46ff230141a72b71",
			},
			expected: expectation{
				count: 50,
				out: &pb.RoomStreamResponse{
					Players: []*pb.Player{{Id: "1", Name: "RexiusTR", Position: &pb.Coordinate{X: 66, Y: 87}, Direction: pb.Direction_STOP}},
				},
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			assert := assert.New(t)
			ctx := context.Background()

			client, closer := newTestServer(ctx)
			defer closer()

			stream, err := client.RoomStream(ctx)
			if !assert.Nil(err) {
				return
			}

			for i := 0; i < tc.expected.count; i++ {
				err = stream.Send(tc.in)
				if !assert.Nil(err) {
					return
				}
			}

			for i := 0; i < tc.expected.count; i++ {
				out, err := stream.Recv()
				if tc.expected.err == nil {
					assert.Nil(err)
					assert.Equal(tc.expected.out.String(), out.String())
				} else {
					assert.Nil(out)
					assert.Equal(tc.expected.err, err)
				}
			}
		})
	}
}
