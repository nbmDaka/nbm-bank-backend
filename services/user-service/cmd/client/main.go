package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nbmDaka/nbm-bank-backend/services/user-service/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)
	ctx = metadata.AppendToOutgoingContext(
		ctx,
		"authorization",
		"Bearer test-token",
	)

	defer cancel()

	response, err := client.GetUser(
		ctx,
		&pb.GetUserRequest{
			Id: 1,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		response,
	)
}
