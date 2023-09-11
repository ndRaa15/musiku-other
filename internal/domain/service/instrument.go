package service

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
)

type InstrumentServiceImpl interface {
	GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error)
	GetByID(ctx context.Context, id uint) (*entity.Instrument, error)
	RentInstrument(ctx context.Context, id uint) (*entity.Instrument, error)
	GetProvince(ctx context.Context, idProvince string) (interface{}, error)
	GetCity(ctx context.Context, idProvince, idCity string) (interface{}, error)
	GetCost(ctx context.Context, cityOrigin, weight, courier string) (interface{}, error)
}
