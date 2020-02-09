package api

import (
	"github.com/go-chi/chi"
)

func (a *API) RegisterRoutes(r *chi.Mux) {
	r.Get("/public-feed", a.GetPublicFeedHandler())
}
