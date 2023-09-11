package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
)

type StudioRepositoryImpl interface {
	GetByID(ctx context.Context, id uint) (*entity.Studio, error)
}
