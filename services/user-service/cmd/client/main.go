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
		"Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpZ1Y2N3hILXRBSEZsRldweWJjNlI3WUFVOWdnV1laUVY4V2VDNXdMS0FrIn0.eyJleHAiOjE3ODg3NjYyODAsImlhdCI6MTc4ODc2NTk4MCwianRpIjoib25ydHJvOjFlNWM5OWNiLTYxMmItZDQxZC0yZjc1LTllZWFhMmE0ZjQxMyIsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9yZWFsbXMvbmJtLWJhbmsiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZDVkZTc5Y2UtZGQxZi00MTAxLTg2NjgtZDI3ODM3NTJjMjMwIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibmJtLWJhY2tlbmQiLCJzaWQiOiIzNjM5MGZjNy03MGZiLTQ1ZGItYjA5MS1mMDE0MmVhNTc2MmIiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIi8qIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIiwiZGVmYXVsdC1yb2xlcy1uYm0tYmFuayJdfSwicmVzb3VyY2VfYWNjZXNzIjp7Im5ibS1iYWNrZW5kIjp7InJvbGVzIjpbIkNVU1RPTUVSIl19LCJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6Ik51cmRhdWxldCBCZWtldG92IiwicHJlZmVycmVkX3VzZXJuYW1lIjoidGVzdHVzZXIiLCJnaXZlbl9uYW1lIjoiTnVyZGF1bGV0IiwiZmFtaWx5X25hbWUiOiJCZWtldG92IiwiZW1haWwiOiJudXJkYXVsZXQuYmVrZXRvdi4yMDA1QGdtYWlsLmNvbSJ9.HeUrDCFTEqo-BzrJsJAxTeeYdEoric8ou7y0A-myX6UwtUOQztE447R_6D8RW4tIAm2qMbnf25n8xgU88tKJYWP5IYPSCqlzRE5_7RtX2wFS3ItCSdkNY8YEeVooOr_uNazS9TAEfPCSxdgXxIQx0fXw20AESoHJB6DI3wr7fo-o1LB549zb_4Tmzt5bnQOCbJ8m9TO_DQ0O0DtN7wK_HFln5OXe0DePXROrDnax3QKwmDdKUpjxvpNuLlztg4B_k0H3Ljnmji1Cx9PZ_nriJV3QUENFpoA5CeQWZHCzN8ty6p-ZsNHjYxgrsI6MRfRyq7TrTnGLmz6bJJ1ENg3o8Q",
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
