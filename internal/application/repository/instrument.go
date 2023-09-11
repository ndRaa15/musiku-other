package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
)

type InstrumentRepository struct {
	db *mysql.DB
}

func NewInstrumentRepository(db *mysql.DB) *InstrumentRepository {
	return &InstrumentRepository{db}
}

func (ir *InstrumentRepository) GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error) {
	var instruments []*entity.Instrument
	if err := ir.db.Debug().WithContext(ctx).Find(&instruments).Error; err != nil {
		return nil, err
	}
	return instruments, nil
}

func (ir *InstrumentRepository) GetByID(ctx context.Context, id uint) (*entity.Instrument, error) {
	var instrument entity.Instrument
	if err := ir.db.Debug().WithContext(ctx).Where("id = ?", id).First(&instrument).Error; err != nil {
		return nil, err
	}
	return &instrument, nil
}

func (ir *InstrumentRepository) Update(ctx context.Context, instrument *entity.Instrument, id uint) (*entity.Instrument, error) {
	if err := ir.db.Debug().WithContext(ctx).Model(&instrument).Where("id = ?", id).Updates(&instrument).Error; err != nil {
		return nil, err
	}
	return instrument, nil
}

func SeedInstrument(db *mysql.DB) error {
	instrument1 := entity.Instrument{
		ID:            1,
		Name:          "Gitar Akustik",
		RentPrice:     75.0,
		Address:       "123 Jalan Musik",
		Description:   "Sebuah gitar akustik berkualitas tinggi.",
		Spesification: "Merek: Yamaha, Warna: Cokelat, Senar: 6, Ukuran: Full",
		Status:        false,
	}

	instrument2 := entity.Instrument{
		ID:            2,
		Name:          "Keyboard Elektronik",
		RentPrice:     120.0,
		Address:       "456 Jalan Harmoni",
		Description:   "Sebuah keyboard elektronik multifungsi.",
		Spesification: "Merek: Roland, Jumlah Tombol: 88, Suara: 300+, Berat: 12 kg",
		Status:        true,
	}

	instrument3 := entity.Instrument{
		ID:            3,
		Name:          "Drum Akustik",
		RentPrice:     100.0,
		Address:       "789 Jalan Rhythm",
		Description:   "Sebuah drum akustik berkualitas tinggi.",
		Spesification: "Merek: Yamaha, Warna: Cokelat, Ukuran: Full",
		Status:        true,
	}

	instrument4 := entity.Instrument{
		ID:            4,
		Name:          "Bass Elektronik",
		RentPrice:     100.0,
		Address:       "789 Jalan Rhythm",
		Description:   "Sebuah bass elektronik berkualitas tinggi.",
		Spesification: "Merek: Yamaha, Warna: Cokelat, Ukuran: Full",
		Status:        false,
	}

	instruments := []entity.Instrument{instrument1, instrument2, instrument3, instrument4}
	for _, instrument := range instruments {
		if err := db.Debug().Create(&instrument).Error; err != nil {
			return err
		}
	}
	return nil
}
