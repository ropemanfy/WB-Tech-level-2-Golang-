package handlers

import (
	"dev11/internal/models"
	"time"
)

type EventStorage interface {
	CreateEvent(event models.Event) error
	UpdateEvent(id string, event models.Event) error
	DeleteEvent(id string) error
	GetEventForDay(date time.Time) ([]models.Event, error)
	GetEventForWeek(date time.Time) ([]models.Event, error)
	GetEventForMonth(date time.Time) ([]models.Event, error)
}
