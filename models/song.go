package models

import (
	"time"

	"github.com/lib/pq"
	"gopkg.in/guregu/null.v3"
)

type Song struct {
	Id            int32          `db:"id"`
	CreatedAt     time.Time      `db:"created_at"`
	Title         string         `db:"title"`
	Artists       pq.StringArray `db:"artists"`
	Album         null.String    `db:"album"`
	SpotifyId     null.String    `db:"spotify_id"`
	AlbumArt      null.String    `db:"album_art"`
	IsrcId        null.String    `db:"isrc_id"`
	AppleMusicId  null.String    `db:"apple_music_id"`
	AppleMusicUrl null.String    `db:"apple_music_url"`
}

type SongWithMeta struct {
	Song
	LikeCount int32 `db:"like_count"`
	IsLiked   bool  `db:"is_liked"`
}
