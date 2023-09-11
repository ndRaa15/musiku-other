package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
)

type InstrumentRepositoryImpl interface {
	GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error)
	GetByID(ctx context.Context, id uint) (*entity.Instrument, error)
	Update(ctx context.Context, instrument *entity.Instrument, id uint) (*entity.Instrument, error)
}
