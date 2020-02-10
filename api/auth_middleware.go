package api

import (
	"context"
	"net/http"

	"github.com/thomasboyt/caroline/models"
)

func GetCurrentUserFromContext(ctx context.Context) *models.User {
	currentUser := ctx.Value("CurrentUser")
	if currentUser != nil {
		return currentUser.(*models.User)
	}
	return nil
}

func (a *API) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *models.User
		token := r.Header.Get("X-Auth-Token")

		ctx := r.Context()
		if token != "" {
			user = a.store.GetUserByAuthToken(token)
			if user != nil {
				ctx = context.WithValue(ctx, "CurrentUser", user)
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
