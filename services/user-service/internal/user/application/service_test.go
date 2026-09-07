package application

import (
	"context"
	"testing"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
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

func TestGetUser(t *testing.T) {

	expectedUser := &domain.User{
		ID:         1,
		KeycloakID: "d5de79ce-dd1f-4101-8668-d2783752c230",
		Email:      "test@test.com",
		FirstName:  "John",
		LastName:   "Smith",
	}

	mockRepo := &mockUserRepository{
		user: expectedUser,
	}

	service := NewUserService(
		mockRepo,
	)

	result, err := service.GetUser(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if result.ID != expectedUser.ID {
		t.Errorf(
			"expected id %d got %d",
			expectedUser.ID,
			result.ID,
		)
	}

	if result.Email != expectedUser.Email {
		t.Errorf(
			"expected email %s got %s",
			expectedUser.Email,
			result.Email,
		)
	}
}

func TestGetUser_UserNotFound(t *testing.T) {

	mockRepo := &mockUserRepository{
		err: domain.ErrUserNotFound,
	}

	service := NewUserService(
		mockRepo,
	)

	_, err := service.GetUser(
		context.Background(),
		999,
	)

	if err == nil {
		t.Fatal(
			"expected error but got nil",
		)
	}

	if err != domain.ErrUserNotFound {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}

func TestGetCurrentUser(t *testing.T) {

	expectedUser := &domain.User{
		ID:         1,
		KeycloakID: "d5de79ce-dd1f-4101-8668-d2783752c230",
		Email:      "test@nbm.com",
		FirstName:  "Daka",
		LastName:   "Developer",
	}

	mockRepo := &mockUserRepository{
		user: expectedUser,
	}

	service := NewUserService(
		mockRepo,
	)

	result, err := service.GetCurrentUser(
		context.Background(),
		"d5de79ce-dd1f-4101-8668-d2783752c230",
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if result.KeycloakID != expectedUser.KeycloakID {
		t.Errorf(
			"expected keycloak id %s got %s",
			expectedUser.KeycloakID,
			result.KeycloakID,
		)
	}

	if result.Email != expectedUser.Email {
		t.Errorf(
			"expected email %s got %s",
			expectedUser.Email,
			result.Email,
		)
	}
}

func TestGetCurrentUser_UserNotFound(t *testing.T) {

	mockRepo := &mockUserRepository{
		err: domain.ErrUserNotFound,
	}

	service := NewUserService(
		mockRepo,
	)

	_, err := service.GetCurrentUser(
		context.Background(),
		"unknown-keycloak-id",
	)

	if err == nil {
		t.Fatal(
			"expected error but got nil",
		)
	}

	if err != domain.ErrUserNotFound {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}
