package service

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/domain/service"
	"github.com/gofrs/uuid"
)

type VenueService struct {
	VenueRepository repository.VenueRepositoryImpl
}

func NewVenueService(VenueRepository repository.VenueRepositoryImpl) service.VenueServiceImpl {
	return &VenueService{VenueRepository}
}

func (vs *VenueService) GetAllVenue(ctx context.Context) ([]*entity.Venue, error) {
	venues, err := vs.VenueRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return venues, nil
}

func (vs *VenueService) GetVenueByID(ctx context.Context, id uint) (*entity.Venue, error) {
	venue, err := vs.VenueRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return venue, nil
}

func (vs *VenueService) RentVenue(ctx context.Context, userID uuid.UUID, venueDayID uint) (*entity.ApplyVenue, error) {
	if err := vs.VenueRepository.UpdateStatusVenueDay(ctx, "PENDING", venueDayID); err != nil {
		return nil, err
	}

	applyVenue := entity.ApplyVenue{
		UserID:     userID,
		VenueDayID: venueDayID,
	}

	res, err := vs.VenueRepository.CreateApplyVenue(ctx, &applyVenue)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vs *VenueService) GetListApplyVenue(ctx context.Context, userID uuid.UUID) ([]*entity.ApplyVenue, error) {
	applyVenues, err := vs.VenueRepository.GetApplyVenueByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return applyVenues, nil
}
