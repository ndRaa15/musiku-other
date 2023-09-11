package service

import (
	"context"
	"os"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/infrastructure/rajaongkir"
)

type InstrumentService struct {
	InstrumentRepository repository.InstrumentRepositoryImpl
}

func NewInstrumentService(instrumentRepository repository.InstrumentRepositoryImpl) *InstrumentService {
	return &InstrumentService{
		InstrumentRepository: instrumentRepository,
	}
}

func (is *InstrumentService) GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error) {
	instruments, err := is.InstrumentRepository.GetAllInstrument(ctx)
	if err != nil {
		return nil, err
	}
	return instruments, nil
}

func (is *InstrumentService) GetByID(ctx context.Context, id uint) (*entity.Instrument, error) {
	instrument, err := is.InstrumentRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return instrument, nil
}

func (is *InstrumentService) RentInstrument(ctx context.Context, id uint) (*entity.Instrument, error) {
	instrument, err := is.InstrumentRepository.Update(ctx, &entity.Instrument{}, id)
	if err != nil {
		return nil, err
	}
	return instrument, nil
}

func (is *InstrumentService) GetProvince(ctx context.Context, idProvince string) (interface{}, error) {
	rajaOngkir := rajaongkir.InitRajaOngkit(os.Getenv("RAJAONGKIR_CREDENTIAL"))
	province, err := rajaOngkir.GetProvince(idProvince)
	if err != nil {
		return nil, err
	}
	return province, nil
}

func (is *InstrumentService) GetCity(ctx context.Context, idProvince, idCity string) (interface{}, error) {
	rajaOngkir := rajaongkir.InitRajaOngkit(os.Getenv("RAJAONGKIR_CREDENTIAL"))
	city, err := rajaOngkir.GetCity(idCity, idProvince)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (is *InstrumentService) GetCost(ctx context.Context, cityOrigin, weight, courier string) (interface{}, error) {
	// need to get the weight first based on the id of the instrument
	// so we just need the city id and courier choose

	rajaOngkir := rajaongkir.InitRajaOngkit(os.Getenv("RAJAONGKIR_CREDENTIAL"))
	cost, err := rajaOngkir.GetCost(cityOrigin, "114", weight, courier)
	if err != nil {
		return nil, err
	}
	return cost, nil
}
