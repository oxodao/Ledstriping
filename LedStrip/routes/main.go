package routes

import (
	"github.com/gorilla/mux"
	"github.com/oxodao/ledstrip/services"
)

func Register(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/state", state(prv))
	r.HandleFunc("/data", data(prv))
	r.HandleFunc("/exec", exec(prv))

	// Favorite
	r.HandleFunc("/favorite/{id}", useFavorite(prv)).Methods("GET")
	r.HandleFunc("/favorite", createFavorite(prv)).Methods("POST")
	r.HandleFunc("/favorite/{id}", deleteFavorite(prv)).Methods("DELETE")
	r.HandleFunc("/favorite/{id}", editFavorite(prv)).Methods("PUT")

	// Colors
	r.HandleFunc("/color/set", setColor(prv))
	r.HandleFunc("/color/spark", spark(prv))

	// Brightness
	r.HandleFunc("/brightness/set", setBrightness(prv))
	r.HandleFunc("/brightness/fade/in", fadeIn(prv))
	r.HandleFunc("/brightness/fade/out", fadeOut(prv))

	// Speed
	r.HandleFunc("/speed/set", setSpeed(prv))

	// Modes
	r.HandleFunc("/mode/set", setMode(prv))
}
