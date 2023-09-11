package entity

import (
	"time"
)

type Studio struct {
	ID           uint        `json:"id" gorm:"autoIncreament;primaryKey"`
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	Description  string      `json:"description"`
	PricePerHour float64     `json:"price_per_hour"`
	OpenHour     string      `json:"open_hour"`
	Status       bool        `json:"status"`
	StartTime    []StartTime `json:"start_time" gorm:"foreignKey:StudioID"`
	EndTime      []EndTime   `json:"end_time" gorm:"foreignKey:StudioID"`
	CreateAt     time.Time   `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt     time.Time   `json:"update_at" gorm:"autoUpdateTime"`
}

type StartTime struct {
	ID          uint   `json:"-" gorm:"autoIncreament;primaryKey"`
	StudioID    uint   `json:"studio_id"`
	StartTime   string `json:"start_time"`
	IsAvailable bool   `json:"isAvailable" gorm:"default:true"`
}

type EndTime struct {
	ID          uint   `json:"-" gorm:"autoIncreament;primaryKey"`
	StudioID    uint   `json:"studio_id"`
	EndTime     string `json:"end_time"`
	IsAvailable bool   `json:"isAvailable" gorm:"default:true"`
}

type RentStudio struct {
	ID        uint    `json:"id" gorm:"autoIncreament;primaryKey"`
	StudioID  uint    `json:"studio_id"`
	UserID    uint    `json:"user_id"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	TotalHour uint    `json:"total_hour"`
	TotalCost float64 `json:"total_cost"`
	Status    string  `json:"status" sql:"type:ENUM('PENDING', 'BOOKED', 'REJECTED')" gorm:"default:'PENDING'"`
}
