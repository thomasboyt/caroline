package api

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	re "github.com/thomasboyt/caroline/resources"
	"github.com/thomasboyt/caroline/services"
)

func parseTimestamp(query url.Values, name string) (*time.Time, *ErrResponse) {
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

func parseTimestamps(query url.Values) (*time.Time, *time.Time, *ErrResponse) {
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

func getCurrentUserId(r *http.Request) int32 {
	var currentUserId int32
	currentUserId = -1
	currentUser := GetCurrentUserFromContext(r.Context())
	if currentUser != nil {
		currentUserId = currentUser.Id
	}
	return currentUserId
}

func (a *API) GetPublicFeedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentUserId := getCurrentUserId(r)
		query := r.URL.Query()

		beforeTimestamp, afterTimestamp, err := parseTimestamps(query)
		if err != nil {
			render.Render(w, r, err)
			return
		}

		feedItems := services.GetPublicFeed(a.store, services.PlaylistArgs{
			BeforeTimestamp: beforeTimestamp,
			AfterTimestamp:  afterTimestamp,
			CurrentUserId:   currentUserId,
		})

		RenderConjson(w, r, feedItems)
	}
}

func (a *API) GetUserPlaylist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentUserId := getCurrentUserId(r)
		userName := chi.URLParam(r, "userName")
		userProfile := services.GetUserProfileByUserName(a.store, userName)

		if userProfile == nil {
			err := CreateErrNotFound("user")
			render.Render(w, r, err)
			return
		}

		query := r.URL.Query()

		beforeTimestamp, afterTimestamp, err := parseTimestamps(query)
		if err != nil {
			render.Render(w, r, *err)
			return
		}

		feedItems := services.GetUserPlaylist(a.store, userProfile.Id, services.PlaylistArgs{
			BeforeTimestamp: beforeTimestamp,
			AfterTimestamp:  afterTimestamp,
			CurrentUserId:   currentUserId,
		})

		resp := re.PlaylistJson{
			Items:       feedItems,
			UserProfile: *userProfile,
		}

		RenderConjson(w, r, resp)
	}
}
