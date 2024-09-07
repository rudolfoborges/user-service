package helper

import (
	"log/slog"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request) error

func (h HandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		errorHandler(w, err)
	}
}

func errorHandler(w http.ResponseWriter, err error) {
	slog.Error("error", "error", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
