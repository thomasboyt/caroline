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

type AggregatedPost struct {
	Timestamp time.Time      `db:"timestamp"`
	UserNames pq.StringArray `db:"user_names"`
	SongId    sql.NullInt32  `db:"song_id"`
	MixtapeId sql.NullInt32  `db:"mixtape_id"`
}
