package routes

import (
	"encoding/json"
	"net/http"

	"github.com/oxodao/ledstrip/services"
)

func state(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := prv.Ledstrip.State.Fetch()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		str, _ := json.Marshal(prv.Ledstrip.State)
		_, _ = w.Write(str)
	}
}
