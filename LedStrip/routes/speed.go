package routes

import (
	"net/http"
	"strconv"

	"github.com/oxodao/ledstrip/services"
)

func setSpeed(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		speed := r.FormValue("speed")
		if len(speed) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		speedInt, err := strconv.ParseUint(speed, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = prv.Ledstrip.Speed.Set(speedInt); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

	}
}
