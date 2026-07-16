package presence

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) UpdateLastSeen(
	ctx context.Context,
	deviceID uuid.UUID,
) error {

	query := `
	UPDATE devices
	SET last_seen = NOW()
	WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		deviceID,
	)

	return err
}