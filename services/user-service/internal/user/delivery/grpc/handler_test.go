package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/auth"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/application"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
	pb "github.com/nbmDaka/nbm-bank-backend/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockUserRepository struct {
	user *domain.User
	err  error
}

func (m *mockUserRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.User, error) {
	return m.user, m.err
}

func (m *mockUserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {
	return nil
}

func (m *mockUserRepository) GetByKeycloakID(
	ctx context.Context,
	keycloakID string,
) (*domain.User, error) {
	return m.user, m.err
}

func TestGetUser_Handler_RepositoryError(t *testing.T) {

	repoErr := errors.New("db connection failure")

	mockRepo := &mockUserRepository{
		err: repoErr,
	}

	service := application.NewUserService(mockRepo)
	handler := NewHandler(service)

	_, err := handler.GetUser(
		context.Background(),
		&pb.GetUserRequest{Id: 1},
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %v", err)
	}

	if st.Code() != codes.Internal {
		t.Errorf("expected status Internal, got %v", st.Code())
	}
}

func TestGetUser_Handler_NotFound(t *testing.T) {

	mockRepo := &mockUserRepository{
		err: domain.ErrUserNotFound,
	}

	service := application.NewUserService(mockRepo)
	handler := NewHandler(service)

	_, err := handler.GetUser(
		context.Background(),
		&pb.GetUserRequest{Id: 999},
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %v", err)
	}

	if st.Code() != codes.NotFound {
		t.Errorf("expected status NotFound, got %v", st.Code())
	}
}

func TestGetCurrentUser_Handler_RepositoryError(t *testing.T) {

	repoErr := errors.New("db connection failure")

	mockRepo := &mockUserRepository{
		err: repoErr,
	}

	service := application.NewUserService(mockRepo)
	handler := NewHandler(service)

	ctx := auth.SetUserID(context.Background(), "test-keycloak-id")

	_, err := handler.GetCurrentUser(
		ctx,
		&pb.GetCurrentUserRequest{},
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %v", err)
	}

	if st.Code() != codes.Internal {
		t.Errorf("expected status Internal, got %v", st.Code())
	}
}

func TestGetCurrentUser_Handler_Unauthenticated(t *testing.T) {

	mockRepo := &mockUserRepository{}

	service := application.NewUserService(mockRepo)
	handler := NewHandler(service)

	// Context without user ID
	_, err := handler.GetCurrentUser(
		context.Background(),
		&pb.GetCurrentUserRequest{},
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %v", err)
	}

	if st.Code() != codes.Unauthenticated {
		t.Errorf("expected status Unauthenticated, got %v", st.Code())
	}
}
