package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/nbmDaka/nbm-bank-backend/proto/user"
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
		5*time.Second,
	)

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpZ1Y2N3hILXRBSEZsRldweWJjNlI3WUFVOWdnV1laUVY4V2VDNXdMS0FrIn0.eyJleHAiOjE3ODg3NzYxMTcsImlhdCI6MTc4ODc3NTgxNywianRpIjoib25ydHJvOjM3NTUyNjg5LWNhNDYtOGUxMC1jN2QyLTM4NGVjYzQxYjcyZSIsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9yZWFsbXMvbmJtLWJhbmsiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZDVkZTc5Y2UtZGQxZi00MTAxLTg2NjgtZDI3ODM3NTJjMjMwIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibmJtLWJhY2tlbmQiLCJzaWQiOiI5ZTgyODRkMy03NzI1LTRjM2EtYjM0Yi02ZDU5NGRlYTAyNTgiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIi8qIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIiwiZGVmYXVsdC1yb2xlcy1uYm0tYmFuayJdfSwicmVzb3VyY2VfYWNjZXNzIjp7Im5ibS1iYWNrZW5kIjp7InJvbGVzIjpbIkNVU1RPTUVSIl19LCJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6Ik51cmRhdWxldCBCZWtldG92IiwicHJlZmVycmVkX3VzZXJuYW1lIjoidGVzdHVzZXIiLCJnaXZlbl9uYW1lIjoiTnVyZGF1bGV0IiwiZmFtaWx5X25hbWUiOiJCZWtldG92IiwiZW1haWwiOiJudXJkYXVsZXQuYmVrZXRvdi4yMDA1QGdtYWlsLmNvbSJ9.cSNJmAH8DCG7wBaUHt4xqf13djj_mSh6GwrBEK5p55ysgHf55DI9qp6n5MA2GzZVOMAnlgMUYA8INBvTbZ--sQQ_JwlovLaSd3Dn_TX7BEDCzRiqU7h-pZtqe_O44ooCwCS0vQl8Kte32QHD_w357dmyfC8wx0KK_5hmdlIeHjv4RIpiuI-EryaeBvU57J3uf1oy_VFSriyqZovq4m20-q6SXSWzOM2omM8LGovxiGyme1ep_1OkKvDO2b-lXgoOZjD5L8-5rpEQqW5_mPN_YuHK9Wd6VSCP2ON8neZ73v4il_Wz0zba_qbQPNq-DY1oqdcvAmxgLd0A_Rl5TFODQw")

	defer cancel()

	resp, err := client.CreateUser(
		ctx,
		&pb.CreateUserRequest{

			KeycloakId: "b37ad3cd-7cbe-493a-9550-97a2466a4e9a",

			Email: "grpc-test@test.com",

			FirstName: "Grpc",

			LastName: "Test",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.User)
}
