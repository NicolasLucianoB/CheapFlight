package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	FlightNumber  string    `json:"flight_number"`
	Departure     string    `json:"departure"`
	Arrival       string    `json:"arrival"`
	Price         float64   `json:"price"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
}
