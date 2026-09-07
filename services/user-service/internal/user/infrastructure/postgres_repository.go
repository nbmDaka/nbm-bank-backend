package infrastructure

import (
	"context"
	"database/sql"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/platform/database"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
)

type PostgresUserRepository struct {
	db database.DB
}

func NewPostgresUserRepository(
	db database.DB,
) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.User, error) {

	query := `
		SELECT
			id,
			email,
			first_name,
			last_name
		FROM users
		WHERE id = $1
	`

	var user domain.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {

	query := `
		INSERT INTO users (
			keycloak_id,
			email,
			first_name,
			last_name
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		user.KeycloakID,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func (r *PostgresUserRepository) GetByKeycloakID(
	ctx context.Context,
	keycloakID string,
) (*domain.User, error) {

	query := `
		SELECT
			id,
			keycloak_id,
			email,
			first_name,
			last_name
		FROM users
		WHERE keycloak_id = $1
	`

	var user domain.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		keycloakID,
	).Scan(
		&user.ID,
		&user.KeycloakID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}
