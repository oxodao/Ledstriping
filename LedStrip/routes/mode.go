package routes

import (
	"github.com/oxodao/ledstrip/services"
	"net/http"
	"strconv"
)

func setMode(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mode := r.FormValue("mode")
		if len(mode) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !prv.ExecuteCommandBoolean("m " + mode) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		modeInt, err := strconv.ParseUint(mode, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		prv.CurrentState.Brightness = modeInt
	}
}
