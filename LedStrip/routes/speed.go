package routes

import (
	"github.com/oxodao/ledstrip/services"
	"net/http"
	"strconv"
)

func setSpeed(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		speed := r.FormValue("speed")
		if len(speed) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !prv.ExecuteCommandBoolean("s " + speed) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		speedInt, err := strconv.ParseUint(speed, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		prv.CurrentState.Brightness = speedInt
	}
}
