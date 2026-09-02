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