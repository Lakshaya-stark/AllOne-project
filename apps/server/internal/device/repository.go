package device

import (
	"context"

	"allone/server/internal/models"

	"github.com/google/uuid"
)

type Repository interface {

	Register(
		ctx context.Context,
		device *models.Device,
	) error

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Device, error)

	ListByUser(
		ctx context.Context,
		userID uuid.UUID,
	) ([]models.Device, error)
}