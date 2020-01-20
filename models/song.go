package models

import "time"

import "database/sql"

import "github.com/lib/pq"

type Song struct {
	Id            int32          `db:"id"`
	CreatedAt     time.Time      `db:"created_at"`
	Title         string         `db:"title"`
	Artists       pq.StringArray `db:"artists"`
	Album         sql.NullString `db:"album"`
	SpotifyId     sql.NullString `db:"spotify_id"`
	AlbumArt      sql.NullString `db:"album_art"`
	IsrcId        sql.NullString `db:"isrc_id"`
	AppleMusicId  sql.NullString `db:"apple_music_id"`
	AppleMusicUrl sql.NullString `db:"apple_music_url"`
}

type SongWithMeta struct {
	Song
	LikeCount int32 `db:"like_count"`
	IsLiked   bool  `db:"is_liked"`
}
