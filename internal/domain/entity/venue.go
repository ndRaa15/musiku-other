package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Venue struct {
	ID          uint       `json:"id" gorm:"autoIncreament;primaryKey"`
	Name        string     `json:"name"`
	Address     string     `json:"username"`
	Description string     `json:"description"`
	VenueDays   []VenueDay `json:"venue_days" gorm:"foreignKey:VenueID"`
	CreateAt    time.Time  `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt    time.Time  `json:"update_at" gorm:"autoUpdateTime"`
}

type VenueDay struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	VenueID   uint    `json:"-"`
	DayID     uint    `json:"-"`
	Day       Day     `json:"day" gorm:"foreignKey:DayID"`
	Salary    float64 `json:"salary"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	Status    string  `json:"status" sql:"type:ENUM('PENDING', 'BOOKED', 'AVAILABLE')" gorm:"default:'AVAILABLE'"`
}
type ApplyVenue struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uuid.UUID `json:"user_id"`
	VenueDayID uint      `json:"-"`
	VenueDay   VenueDay  `json:"venue_day" gorm:"foreignKey:VenueDayID"`
}
