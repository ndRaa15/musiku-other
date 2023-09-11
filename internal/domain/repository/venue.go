package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/gofrs/uuid"
)

type VenueRepositoryImpl interface {
	GetAll(ctx context.Context) ([]*entity.Venue, error)
	GetByID(ctx context.Context, id uint) (*entity.Venue, error)
	UpdateStatusVenueDay(ctx context.Context, status string, venueDayID uint) error
	CreateApplyVenue(ctx context.Context, applyVenue *entity.ApplyVenue) (*entity.ApplyVenue, error)
	GetApplyVenueByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.ApplyVenue, error)
}
