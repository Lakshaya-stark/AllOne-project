package auth

import (
	"context"

	"allone/server/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) CreateUser(
	ctx context.Context,
	user *models.User,
) error {

	query := `
INSERT INTO users
(id, username, email, password_hash, created_at, updated_at)
VALUES ($1,$2,$3,$4,$5,$6)
`

	_, err := r.db.Exec(
		ctx,
		query,
		user.ID,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *PostgresRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {

	var user models.User

	query := `
SELECT
	id,
	username,
	email,
	password_hash,
	created_at,
	updated_at
FROM users
WHERE email=$1
`

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}