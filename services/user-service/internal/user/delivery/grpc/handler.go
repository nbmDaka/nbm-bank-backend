package grpc


import (
	"context"

	pb "github.com/nbmDaka/nbm-bank-backend/services/user-service/proto/user"

)


type Handler struct {

	pb.UnimplementedUserServiceServer

}



func NewHandler() *Handler {

	return &Handler{}

}



func (h *Handler) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
)(
	*pb.GetUserResponse,
	error,
){


	return &pb.GetUserResponse{

		Id: req.Id,

		Email: "test@bank.com",

		FirstName: "John",

		LastName: "Doe",

	},nil

}