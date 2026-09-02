package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/nbmDaka/nbm-bank-backend/services/user-service/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {


	conn, err := grpc.NewClient(
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



	fmt.Println("User:")
	fmt.Println("ID:", response.Id)
	fmt.Println("Email:", response.Email)
	fmt.Println("First name:", response.FirstName)
	fmt.Println("Last name:", response.LastName)

}