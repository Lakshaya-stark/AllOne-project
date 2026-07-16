package presence

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {

	SetOnline(
		ctx context.Context,
		deviceID uuid.UUID,
	) error

	SetOffline(
		ctx context.Context,
		deviceID uuid.UUID,
	) error

	UpdateLastSeen(
		ctx context.Context,
		deviceID uuid.UUID,
	) error

	IsOnline(
		ctx context.Context,
		deviceID uuid.UUID,
	) (bool, error)
}