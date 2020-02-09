package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	Id        int32         `db:"id"`
	CreatedAt time.Time     `db:"created_at"`
	UserId    int32         `db:"user_id"`
	SongId    sql.NullInt32 `db:"song_id"`
	MixtapeId sql.NullInt32 `db:"mixtape_id"`
}

type PostWithConnections interface {
	GetSongId() sql.NullInt32
	GetMixtapeId() sql.NullInt32
}

type UserPost struct {
	Timestamp time.Time     `db:"timestamp"`
	SongId    sql.NullInt32 `db:"song_id"`
	MixtapeId sql.NullInt32 `db:"mixtape_id"`
}

func (p UserPost) GetSongId() sql.NullInt32 {
	return p.SongId
}

func (p UserPost) GetMixtapeId() sql.NullInt32 {
	return p.MixtapeId
}

type AggregatedPost struct {
	Timestamp time.Time      `db:"timestamp"`
	UserNames pq.StringArray `db:"user_names"`
	SongId    sql.NullInt32  `db:"song_id"`
	MixtapeId sql.NullInt32  `db:"mixtape_id"`
}

func (p AggregatedPost) GetSongId() sql.NullInt32 {
	return p.SongId
}

func (p AggregatedPost) GetMixtapeId() sql.NullInt32 {
	return p.MixtapeId
}
