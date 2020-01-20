package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/thomasboyt/jam-buds-golang/api"
	"github.com/thomasboyt/jam-buds-golang/store"
)

func main() {
	log.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal("failed loading .env: ", err)
		}
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	store := store.New()
	a := api.New(store)
	a.RegisterRoutes(r)

	log.Println("Server starting on port 3333")

	err = http.ListenAndServe(":3333", r)

	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
