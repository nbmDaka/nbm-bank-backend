package infrastructure

import (
	"context"
	"database/sql"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(
	db *sql.DB,
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
			password_hash,
			first_name,
			last_name,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	user := &domain.User{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {

	query := `
		INSERT INTO users (
			email,
			password_hash,
			first_name,
			last_name
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
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