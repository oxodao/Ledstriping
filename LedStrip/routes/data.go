package routes

import (
	"encoding/json"
	"net/http"

	"github.com/oxodao/ledstrip/models"
	"github.com/oxodao/ledstrip/services"
)

func data(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := models.Data{
			Modes:     prv.Ledstrip.Mode.GetAvailableModes(),
			Favorites: prv.Config.Favorites,
		}

		d, _ := json.Marshal(data)
		w.Write(d)
	}
}
