package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/application"
	pb "github.com/nbmDaka/nbm-bank-backend/proto/user"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/mapper"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/validation"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/auth"
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

func (h *Handler) GetCurrentUser(
	ctx context.Context,
	req *pb.GetCurrentUserRequest,
) (*pb.GetUserResponse,error){


	keycloakID, ok := auth.GetUserID(ctx)

	if !ok {
		return nil,status.Error(
			codes.Unauthenticated,
			"user identity missing",
		)
	}


	user, err := h.service.GetCurrentUser(
		ctx,
		keycloakID,
	)


	if err != nil {

		if errors.Is(err, domain.ErrUserNotFound) {
			return nil,status.Error(
				codes.NotFound,
				"user not found",
			)
		}


		return nil,status.Error(
			codes.Internal,
			"internal server error",
		)
	}


	return &pb.GetUserResponse{
		User: mapper.ToProto(user),
	},nil
}

func (h *Handler) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {


	user := &domain.User{

		KeycloakID: req.KeycloakId,

		Email: req.Email,

		FirstName: req.FirstName,

		LastName: req.LastName,
	}


	err := h.service.CreateUserProfile(
		ctx,
		user,
	)


	if err != nil {

		return nil, status.Error(
			codes.Internal,
			err.Error(),
		)
	}


	return &pb.CreateUserResponse{
		User: mapper.ToProto(user),
	}, nil
}