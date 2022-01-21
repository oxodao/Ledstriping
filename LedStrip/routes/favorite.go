package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oxodao/ledstrip/models"
	"github.com/oxodao/ledstrip/services"

	"github.com/google/uuid"
)

func useFavorite(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		found := false
		for _, f := range prv.Config.Favorites {
			if vars["id"] != f.ID {
				continue
			}

			if err := prv.Ledstrip.SetState(f.Color, f.Brightness, f.Mode, f.Speed); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			found = true
		}

		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

func createFavorite(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var favorite models.Favorite
		err := json.NewDecoder(r.Body).Decode(&favorite)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		favorite.ID = uuid.NewString()

		for _, f := range prv.Config.Favorites {
			if favorite.ID == f.ID {
				w.WriteHeader(http.StatusConflict)
				return
			}
		}

		prv.Config.Favorites = append(prv.Config.Favorites, favorite)

		err = prv.Config.Save()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		str, _ := json.Marshal(favorite)
		_, _ = w.Write(str)
	}
}

func deleteFavorite(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		for i, f := range prv.Config.Favorites {
			if vars["id"] != f.ID {
				continue
			}

			prv.Config.Favorites = append(prv.Config.Favorites[:i], prv.Config.Favorites[i+1:]...)
			break
		}

		err := prv.Config.Save()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusGone)
	}
}

func editFavorite(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		for _, f := range prv.Config.Favorites {
			if vars["id"] != f.ID {
				continue
			}

			var favorite *models.Favorite = &f
			err := json.NewDecoder(r.Body).Decode(favorite)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			favorite.ID = vars["id"]
			break
		}

		err := prv.Config.Save()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
