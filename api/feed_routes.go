package api

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/render"
	"github.com/thomasboyt/jam-buds-golang/services"
)

func parseTimestamp(query url.Values, name string) (*time.Time, *ErrInvalidParameter) {
	value := query.Get(name)
	if value == "" {
		return nil, nil
	}
	timestamp, err := time.Parse(time.RFC3339, value)
	if err != nil {
		err := CreateErrInvalidParameter(name, "could not be parsed as RFC3339 time")
		return nil, &err
	}
	return &timestamp, nil
}

func parseTimestamps(query url.Values) (*time.Time, *time.Time, *ErrInvalidParameter) {
	beforeTimestamp, err := parseTimestamp(query, "before")
	if err != nil {
		return nil, nil, err
	}
	afterTimestamp, err := parseTimestamp(query, "after")
	if err != nil {
		return nil, nil, err
	}
	return beforeTimestamp, afterTimestamp, nil
}

func (a *API) GetPublicFeedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		beforeTimestamp, afterTimestamp, err := parseTimestamps(query)
		if err != nil {
			render.Render(w, r, err)
			return
		}

		feedItems := services.GetPublicFeed(a.store, beforeTimestamp, afterTimestamp)

		RenderConjson(w, r, feedItems)
	}
}
