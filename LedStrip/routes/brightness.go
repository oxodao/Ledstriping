package routes

import (
	"net/http"
	"strconv"

	"github.com/oxodao/ledstrip/services"
)

func setBrightness(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brightness := r.FormValue("brightness")
		if len(brightness) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		brightnessInt, err := strconv.ParseUint(brightness, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = prv.Ledstrip.Brightness.Set(brightnessInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func fadeIn(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if prv.Ledstrip.Brightness.FadeIn() != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func fadeOut(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if prv.Ledstrip.Brightness.FadeOut() != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
