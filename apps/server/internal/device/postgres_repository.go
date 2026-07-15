package device

import (
	"context"

	"allone/server/internal/models"

	"github.com/google/uuid"
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

func (r *PostgresRepository) Register(
	ctx context.Context,
	device *models.Device,
) error {

	query := `
	INSERT INTO devices
	(
		id,
		user_id,
		name,
		platform,
		device_type,
		public_key,
		last_seen,
		created_at
	)
	VALUES
	($1,$2,$3,$4,$5,$6,$7,$8)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		device.ID,
		device.UserID,
		device.Name,
		device.Platform,
		device.DeviceType,
		device.PublicKey,
		device.LastSeen,
		device.CreatedAt,
	)

	return err
}
func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Device, error){

	var device models.Device

	query := `
	SELECT
		id,
		user_id,
		name,
		platform,
		device_type,
		public_key,
		last_seen,
		created_at
	FROM devices
	WHERE id = $1
	`

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&device.ID,
		&device.UserID,
		&device.Name,
		&device.Platform,
		&device.DeviceType,
		&device.PublicKey,
		&device.LastSeen,
		&device.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *PostgresRepository) ListByUser(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Device, error) {

	query := `
	SELECT
		id,
		user_id,
		name,
		platform,
		device_type,
		public_key,
		last_seen,
		created_at
	FROM devices
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(
		ctx,
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var devices []models.Device

	for rows.Next() {

		var device models.Device

		err := rows.Scan(
			&device.ID,
			&device.UserID,
			&device.Name,
			&device.Platform,
			&device.DeviceType,
			&device.PublicKey,
			&device.LastSeen,
			&device.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}