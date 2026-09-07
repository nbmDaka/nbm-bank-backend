package domain

import "time"

type User struct {
	ID int64

	KeycloakID string

	Email string

	PasswordHash string

	FirstName string

	LastName string

	CreatedAt time.Time

	UpdatedAt time.Time
}