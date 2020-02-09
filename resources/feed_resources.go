package resources

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type SongJson struct {
	Id            int32
	Title         string
	Artists       []string
	IsLiked       bool
	LikeCount     int32
	Album         null.String
	AlbumArt      null.String
	SpotifyId     null.String
	AppleMusicId  null.String
	AppleMusicUrl null.String
}

type MixtapePreviewJson struct {
	Id          int32
	CreatedAt   time.Time
	Title       string
	Slug        string
	UserId      int32
	PublishedAt time.Time
	SongCount   int32
	AuthorName  string
}

type FeedItemJson struct {
	Timestamp time.Time
	UserNames []string
	Song      *SongJson
	Mixtape   *MixtapePreviewJson
}
