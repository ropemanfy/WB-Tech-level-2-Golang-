package handlers

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		logrus.Printf("method: %s  URI: %s  lead time: %s", r.Method, r.RequestURI, time.Since(start))
	})
}
