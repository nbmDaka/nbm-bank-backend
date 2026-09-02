package grpc

import (
	"context"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/application"
	pb "github.com/nbmDaka/nbm-bank-backend/services/user-service/proto/user"
)


type Handler struct {
	service *application.UserService

	pb.UnimplementedUserServiceServer
}

func NewHandler(
	service *application.UserService,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse,error){

	user,err := h.service.GetUser(
		ctx,
		req.Id,
	)

	if err != nil {
		return nil,err
	}


	return &pb.GetUserResponse{

		User:&pb.User{
			Id:user.ID,
			Email:user.Email,
			FirstName:user.FirstName,
			LastName:user.LastName,
		},

	},nil
}