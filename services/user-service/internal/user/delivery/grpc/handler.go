package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/application"
	pb "github.com/nbmDaka/nbm-bank-backend/services/user-service/proto/user"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/mapper"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/validation"
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

	err := validation.ValidateUserID(
		req.Id,
	)

	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Invalid user ID",
		)
	}

	user,err := h.service.GetUser(
		ctx,
		req.Id,
	)

	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil,status.Error(
				codes.NotFound,
				"User not found",
			)
		}


		return nil,status.Error(
			codes.Internal,
			"Internal server error",
		)
	}


	return &pb.GetUserResponse{
		User: mapper.ToProto(user),
	}, nil
}