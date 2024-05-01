package handlers

import (
	"dev11/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	EventStorage
}

func NewHandler(storage EventStorage) *Handler {
	return &Handler{storage}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.eventForDay)
	mux.HandleFunc("/events_for_week", h.eventForWeek)
	mux.HandleFunc("/events_for_month", h.eventForMonth)

	handler := Logging(mux)
	return handler
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err400(w)
		return
	}

	event, err := decodeJson(r)
	if err != nil {
		logrus.Println(err)
		err400(w)
		return
	}

	err = h.CreateEvent(event)
	if err != nil {
		logrus.Println(err)
		err503(w)
		return
	}

	respPost(w)
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err400(w)
		return
	}

	event, err := decodeJson(r)
	if err != nil {
		logrus.Println(err)
		err400(w)
		return
	}

	id, err := getId(r.URL)
	if err != nil {
		logrus.Println(err)
		err400(w)
		return
	}

	err = h.UpdateEvent(id, event)
	if err != nil {
		logrus.Println(err)
		err503(w)
		return
	}

	respPost(w)
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err400(w)
		return
	}

	id, err := getId(r.URL)
	if err != nil {
		logrus.Println(err)
		err400(w)
		return
	}

	err = h.DeleteEvent(id)
	if err != nil {
		logrus.Println(err)
		err503(w)
		return
	}

	respPost(w)
}

func (h *Handler) eventForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err400(w)
		return
	}

	date, err := getDate(r.URL)
	if err != nil {
		logrus.Println(err)
		err400(w)
		return
	}

	events, err := h.GetEventForDay(date)
	if err != nil {
		logrus.Println(err)
		err503(w)
		return
	}

	respGet(w, events)
}

func (h *Handler) eventForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err400(w)
		return
	}

	date, err := getDate(r.URL)
	if err != nil {
		logrus.Println(err)
		err400(w)
		return
	}

	events, err := h.GetEventForWeek(date)
	if err != nil {
		logrus.Println(err)
		err503(w)
		return
	}

	respGet(w, events)
}

func (h *Handler) eventForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err400(w)
		return
	}

	date, err := getDate(r.URL)
	if err != nil {
		logrus.Println(err)
		err400(w)
		return
	}

	events, err := h.GetEventForMonth(date)
	if err != nil {
		logrus.Println(err)
		err503(w)
		return
	}

	respGet(w, events)
}

func decodeJson(r *http.Request) (models.Event, error) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return models.Event{}, fmt.Errorf("decoder: %v", err)
	}
	return event, nil
}

func validate(date string) (time.Time, error) {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time, fmt.Errorf("invalid date")
	}
	return time, nil
}

func getDate(url *url.URL) (time.Time, error) {
	date := url.Query().Get("date")

	time, err := validate(date)
	if err != nil {
		return time, err
	}

	return time, nil
}

func getId(url *url.URL) (string, error) {
	id := url.Query().Get("id")

	if len(id) == 0 {
		return "", fmt.Errorf("epmty id")
	}

	return id, nil
}

func err400(w http.ResponseWriter) {
	resp, _ := json.Marshal(map[string]int{
		"error": http.StatusBadRequest,
	})
	w.Write(resp)
}

func err503(w http.ResponseWriter) {
	resp, _ := json.Marshal(map[string]int{
		"error": http.StatusServiceUnavailable,
	})
	w.Write(resp)
}

func respPost(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, _ := json.Marshal(map[string]string{
		"result": "successfully",
	})
	w.Write(result)
}

func respGet(w http.ResponseWriter, events []models.Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, _ := json.Marshal(map[string][]models.Event{
		"result": events,
	})
	w.Write(result)
}
