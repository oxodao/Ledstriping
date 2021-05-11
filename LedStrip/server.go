package main

import (
	"embed"
	"github.com/oxodao/ledstrip/routes"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oxodao/ledstrip/services"
)

//go:embed build
var frontendFS embed.FS

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func RunServer(prv *services.Provider) {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(jsonMiddleware)
	routes.Register(prv, api)

	r.PathPrefix("/").Handler(http.FileServer(getFrontendFS()))

	srv := &http.Server {
		Handler: r,
		Addr: prv.Config.ListeningUrl,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getFrontendFS() http.FileSystem {
	fsys, err := fs.Sub(frontendFS, "build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}