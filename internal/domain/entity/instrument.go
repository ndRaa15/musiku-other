package entity

import (
	"time"
)

type Instrument struct {
	ID            uint      `json:"id" gorm:"autoIncreament;primaryKey"`
	Name          string    `json:"name"`
	RentPrice     float64   `json:"rent_price"`
	Address       string    `json:"address"`
	Description   string    `json:"description"`
	Spesification string    `json:"spesification"`
	Status        bool      `json:"status"`
	CreateAt      time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt      time.Time `json:"update_at" gorm:"autoUpdateTime"`
}

type RentInstrument struct {
	ID             uint    `json:"id" gorm:"autoIncreament;primaryKey"`
	InstrumentID   uint    `json:"instrument_id"`
	LengthLoan     uint    `json:"length_loan"`
	UserID         uint    `json:"user_id"`
	Cost           float64 `json:"cost"`
	TotalCost      float64 `json:"total_cost"`
	DayRent        string  `json:"day_rent"`
	EstimationTime string  `json:"estimation_time"`
}
