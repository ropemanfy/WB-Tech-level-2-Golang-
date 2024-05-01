package storage

import (
	"dev11/internal/models"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type EventStorage interface {
	CreateEvent(event models.Event) error
	UpdateEvent(id string, event models.Event) error
	DeleteEvent(id string) error
	GetEventForDay(date time.Time) ([]models.Event, error)
	GetEventForWeek(date time.Time) ([]models.Event, error)
	GetEventForMonth(date time.Time) ([]models.Event, error)
}

type storage struct {
	data map[string]models.Event
	sync.RWMutex
}

func NewStorage() EventStorage {
	return &storage{data: make(map[string]models.Event)}
}

func (s *storage) CreateEvent(event models.Event) error {
	s.Lock()
	defer s.Unlock()

	for _, v := range s.data {
		if event == v {
			return fmt.Errorf("event already exists")
		}
	}

	id := uuid.New().String()
	s.data[id] = event
	return nil
}

func (s *storage) UpdateEvent(id string, event models.Event) error {
	s.Lock()
	defer s.Unlock()

	_, ok := s.data[id]
	if !ok {
		return fmt.Errorf("event does not exist")
	}

	s.data[id] = event
	return nil
}

func (s *storage) DeleteEvent(id string) error {
	s.Lock()
	defer s.Unlock()

	_, ok := s.data[id]
	if !ok {
		return fmt.Errorf("event does not exist")
	}

	delete(s.data, id)
	return nil
}

func (s *storage) GetEventForDay(date time.Time) ([]models.Event, error) {
	var result []models.Event
	s.Lock()
	defer s.RUnlock()
	for _, v := range s.data {
		if v.Date == date {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return result, fmt.Errorf("no events on the selected date")
	}

	return result, nil
}

func (s *storage) GetEventForWeek(date time.Time) ([]models.Event, error) {
	var result []models.Event
	s.Lock()
	defer s.RUnlock()

	week := date.AddDate(0, 0, 7)
	for _, event := range s.data {
		after := event.Date.After(date)
		before := event.Date.Before(week)
		if after && before {
			result = append(result, event)
		}
	}

	if len(result) == 0 {
		return result, fmt.Errorf("no events on the selected date")
	}

	return result, nil
}

func (s *storage) GetEventForMonth(date time.Time) ([]models.Event, error) {
	var result []models.Event
	s.Lock()
	defer s.RUnlock()

	week := date.AddDate(0, 1, 0)
	for _, event := range s.data {
		after := event.Date.After(date)
		before := event.Date.Before(week)
		if after && before {
			result = append(result, event)
		}
	}

	if len(result) == 0 {
		return result, fmt.Errorf("no events on the selected date")
	}

	return result, nil
}
