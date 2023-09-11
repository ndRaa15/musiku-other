package service

import (
	"context"
	"errors"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
)

type StudioService struct {
	StudioRepository repository.StudioRepositoryImpl
}

func NewStudioService(studioRepository repository.StudioRepositoryImpl) *StudioService {
	return &StudioService{
		StudioRepository: studioRepository,
	}
}

func (ss *StudioService) GetAllStudio(ctx context.Context) ([]*entity.Studio, error) {
	return nil, nil
}

func (ss *StudioService) GetStudioByID(ctx context.Context) {

}

func (ss *StudioService) RentStudio(ctx context.Context, studioID uint, startTime, endTime string) (*entity.RentStudio, error) {
	//Using for loop
	studio, err := ss.StudioRepository.GetByID(ctx, studioID)
	if err != nil {
		return nil, err
	}

	var indexStartTime int
	var indexEndTime int

	for i := 0; i < len(studio.StartTime); i++ {
		if studio.StartTime[i].StartTime == startTime {
			indexStartTime = i
		}
		if studio.EndTime[i].EndTime == endTime {
			indexEndTime = i
		}
	}

	if indexStartTime > indexEndTime {
		return nil, errors.New("INVALID TIME")
	}

	// Make start time and end time to be slice
	for i := indexStartTime; i <= indexEndTime; i++ {
		studio.StartTime[i].IsAvailable = false
		studio.EndTime[i].IsAvailable = false
	}

	//Implement update studio
	//Create rent studio entity
	return nil, nil
}
