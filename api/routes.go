package api

import (
	"github.com/go-chi/chi"
)

func (a *API) RegisterRoutes(r *chi.Mux) {
	r.Get("/public-feed", a.GetPublicFeedHandler())
	r.Get("/playlists/{userName}", a.GetUserPlaylist())

	// r.Group(func(r chi.Router) {
	// 	r.Use(requireAuthToken)
	// })
}
