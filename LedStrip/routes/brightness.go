package routes

import (
	"github.com/oxodao/ledstrip/services"
	"net/http"
	"strconv"
)

func setBrightness(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brightness := r.FormValue("brightness")
		if len(brightness) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !prv.ExecuteCommandBoolean("b " + brightness) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		brightnessInt, err := strconv.ParseUint(brightness, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		prv.CurrentState.Brightness = brightnessInt
	}
}

func fadeIn(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !prv.ExecuteCommandBoolean("fadein") {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func fadeOut(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !prv.ExecuteCommandBoolean("fadeout") {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}