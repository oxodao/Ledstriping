package routes

import (
	"encoding/json"
	"net/http"

	"github.com/oxodao/ledstrip/services"
)

func exec(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := prv.Ledstrip.ExecuteCommand(r.FormValue("command"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
		}

		jsonned, _ := json.Marshal(struct {
			Response string
		}{
			Response: resp,
		})

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonned)
	}
}
