package routes

import (
	"github.com/oxodao/ledstrip/services"
	"github.com/oxodao/ledstrip/utils"
	"net/http"
)

func setColor(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := r.FormValue("color")
		if len(color) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !prv.ExecuteCommandBoolean("c " + utils.GetBoardFromHex(color)) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		prv.CurrentState.Color = color
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

		_, _ = prv.ExecuteCommand("d " + duration)
		_, _ = prv.ExecuteCommand("sp " + utils.GetBoardFromHex(color))
	}
}
