package api

import "github.com/thomasboyt/jam-buds-golang/store"

type API struct {
	store *store.Store
}

// New instantiates a new API object
func New(store *store.Store) *API {
	return &API{
		store,
	}
}
