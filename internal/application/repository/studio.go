package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
)

type StudioRepository struct {
	db *mysql.DB
}

func NewStudioRepository(db *mysql.DB) *StudioRepository {
	return &StudioRepository{
		db: db,
	}
}

func (sr *StudioRepository) GetAll(ctx context.Context) ([]*entity.Studio, error) {
	var studios []*entity.Studio
	if err := sr.db.Debug().WithContext(ctx).Find(&studios).Error; err != nil {
		return nil, err
	}
	return studios, nil
}

func (sr *StudioRepository) GetByID(ctx context.Context, id uint) (*entity.Studio, error) {
	var studio entity.Studio
	if err := sr.db.Debug().WithContext(ctx).Where("id = ?", id).First(&studio).Error; err != nil {
		return nil, err
	}
	return &studio, nil
}

func (sr *StudioRepository) RentStudio(ctx context.Context, rentStudio *entity.RentStudio) (*entity.RentStudio, error) {
	if err := sr.db.Debug().WithContext(ctx).Create(&rentStudio).Error; err != nil {
		return nil, err
	}
	return rentStudio, nil
}
