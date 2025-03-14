package models

import (
	"time"
)

type Flight struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" gorm:"index"` // Alterado para *time.Time
	FlightNumber  string     `json:"flight_number"`
	Departure     string     `json:"departure"`
	Arrival       string     `json:"arrival"`
	Price         float64    `json:"price"`
	DepartureTime time.Time  `json:"departure_time"`
	ArrivalTime   time.Time  `json:"arrival_time"`
}
