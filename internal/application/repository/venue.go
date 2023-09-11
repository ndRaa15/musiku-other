package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
	"github.com/gofrs/uuid"
)

type VenueRepository struct {
	db *mysql.DB
}

func NewVenueRepository(db *mysql.DB) repository.VenueRepositoryImpl {
	return &VenueRepository{db}
}

func (vr *VenueRepository) GetAll(ctx context.Context) ([]*entity.Venue, error) {
	var venues []*entity.Venue
	if err := vr.db.Debug().WithContext(ctx).Preload("VenueDays.Day").Find(&venues).Error; err != nil {
		return nil, err
	}
	return venues, nil
}

func (vr *VenueRepository) GetByID(ctx context.Context, id uint) (*entity.Venue, error) {
	var venue entity.Venue
	if err := vr.db.Debug().WithContext(ctx).Preload("VenueDays.Day").Where("id = ?", id).First(&venue).Error; err != nil {
		return nil, err
	}
	return &venue, nil
}

func (vr *VenueRepository) UpdateStatusVenueDay(ctx context.Context, status string, venueDayID uint) error {
	var venue entity.Venue
	if err := vr.db.Debug().WithContext(ctx).Model(&venue).Where("id = ?", venueDayID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (vr *VenueRepository) CreateApplyVenue(ctx context.Context, applyVenue *entity.ApplyVenue) (*entity.ApplyVenue, error) {
	if err := vr.db.Debug().WithContext(ctx).Create(&applyVenue).Error; err != nil {
		return nil, err
	}
	return applyVenue, nil
}

func (vr *VenueRepository) GetApplyVenueByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.ApplyVenue, error) {
	var applyVenues []*entity.ApplyVenue
	if err := vr.db.Debug().WithContext(ctx).Where("user_id = ?", userID).Find(&applyVenues).Error; err != nil {
		return nil, err
	}
	return applyVenues, nil
}

func SeedVenue(db *mysql.DB) error {
	var venues []entity.Venue
	var venueDays []entity.VenueDay

	venue1 := entity.Venue{
		ID:          1,
		Name:        "Kopi Soa",
		Address:     "Jl. Soalawas, No. 1",
		Description: "Kopi Soa adalah tempat nongkrong yang asik",
	}

	venue2 := entity.Venue{
		ID:          2,
		Name:        "Cafe Seru",
		Address:     "Jl. Pahlawan, No. 10",
		Description: "Cafe Seru adalah tempat nongkrong yang cozy",
	}

	venue3 := entity.Venue{
		ID:          3,
		Name:        "Warung Nyaman",
		Address:     "Jl. Nyaman, No. 5",
		Description: "Warung Nyaman adalah tempat makan enak dan murah",
	}

	venues = append(venues, venue1, venue2, venue3)

	for _, venue := range venues {
		if err := db.Create(&venue).Error; err != nil {
			return err
		}
	}

	venue1Days := entity.VenueDay{
		VenueID:   1,
		DayID:     1,
		Salary:    50_000,
		StartTime: "10.00",
		EndTime:   "22.00",
	}

	venue2Days := entity.VenueDay{
		VenueID:   1,
		DayID:     1,
		Salary:    50_000,
		StartTime: "10.00",
		EndTime:   "22.00",
	}

	venue3Days := entity.VenueDay{
		VenueID:   1,
		DayID:     1,
		Salary:    50_000,
		StartTime: "10.00",
		EndTime:   "22.00",
	}

	venueDays = append(venueDays, venue1Days, venue2Days, venue3Days)

	for _, venueDay := range venueDays {
		if err := db.Create(&venueDay).Error; err != nil {
			return err
		}
	}
	return nil
}
