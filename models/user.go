package models

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type User struct {
	Id                  int32       `db:"id"`
	CreatedAt           time.Time   `db:"created_at"`
	Name                string      `db:"name"`
	Email               string      `db:"email"`
	TwitterName         null.String `db:"twitter_name"`
	TwitterId           null.String `db:"twitter_id"`
	TwitterToken        null.String `db:"twitter_token"`
	TwitterSecrets      null.String `db:"twitter_secret"`
	SpotifyAccessToken  null.String `db:"spotify_access_token"`
	SpotifyRefreshToken null.String `db:"spotify_refresh_token"`
	SpotifyExpiresAt    null.Time   `db:"spotify_expires_at"`
	ShowInPublicFeed    null.Bool   `db:"show_in_public_feed"`
}
