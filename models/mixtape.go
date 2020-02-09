package models

import (
	"database/sql"
	"time"
)

type Mixtape struct {
	Id          int32        `db:"id"`
	CreatedAt   time.Time    `db:"created_at"`
	Title       string       `db:"title"`
	Slug        string       `db:"slug"`
	UserId      int32        `db:"user_id"`
	PublishedAt sql.NullTime `db:"published_at"`
}

type MixtapePreview struct {
	Mixtape
	SongCount  int32  `db:"song_count"`
	AuthorName string `db:"author_name"`
}
