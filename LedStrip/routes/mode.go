package routes

import (
	"net/http"

	"github.com/oxodao/ledstrip/services"
)

func setMode(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mode := r.FormValue("mode")
		if len(mode) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := prv.Ledstrip.Mode.Set(mode); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
