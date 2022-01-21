package routes

import (
	"net/http"

	"github.com/oxodao/ledstrip/services"
)

func setColor(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := r.FormValue("color")
		if len(color) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if prv.Ledstrip.Color.Set(color) != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func spark(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := r.FormValue("color")
		duration := r.FormValue("duration")
		if len(color) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := prv.Ledstrip.Color.Spark(color, duration)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
