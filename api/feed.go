package api

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/thomasboyt/jam-buds-golang/models"
)

type songJson struct {
	Id            int32
	Title         string
	Artists       []string
	IsLiked       bool
	LikeCount     int32
	Album         *string
	AlbumArt      *string
	SpotifyId     *string
	AppleMusicId  *string
	AppleMusicUrl *string
}

type mixtapePreviewJson struct {
	Id          int32
	CreatedAt   time.Time
	Title       string
	Slug        string
	UserId      int32
	PublishedAt time.Time // we can assume this is always present in the feed
	SongCount   int32
	AuthorName  string
}

type feedItemJson struct {
	Timestamp time.Time
	UserNames []string
	Song      *songJson
	Mixtape   *mixtapePreviewJson
}

func mapSongsById(songs []models.SongWithMeta) map[int32]models.SongWithMeta {
	m := make(map[int32]models.SongWithMeta)
	for _, song := range songs {
		m[song.Id] = song
	}
	return m
}

func mapMixtapesById(mixtapes []models.MixtapePreview) map[int32]models.MixtapePreview {
	m := make(map[int32]models.MixtapePreview)
	for _, mixtape := range mixtapes {
		m[mixtape.Id] = mixtape
	}
	return m
}

func serializeSong(song models.SongWithMeta) songJson {
	songJson := songJson{
		Id:        song.Id,
		Title:     song.Title,
		Artists:   song.Artists,
		IsLiked:   song.IsLiked,
		LikeCount: song.LikeCount,
	}

	// lmao
	if song.Album.Valid {
		songJson.Album = &song.Album.String
	}
	if song.AlbumArt.Valid {
		songJson.AlbumArt = &song.AlbumArt.String
	}
	if song.SpotifyId.Valid {
		songJson.SpotifyId = &song.SpotifyId.String
	}
	if song.AppleMusicId.Valid {
		songJson.AppleMusicId = &song.SpotifyId.String
	}
	if song.AppleMusicUrl.Valid {
		songJson.AppleMusicUrl = &song.SpotifyId.String
	}

	return songJson
}

func serializeMixtape(mixtape models.MixtapePreview) mixtapePreviewJson {
	return mixtapePreviewJson{
		Id:          mixtape.Id,
		CreatedAt:   mixtape.CreatedAt,
		Title:       mixtape.Title,
		Slug:        mixtape.Slug,
		UserId:      mixtape.UserId,
		PublishedAt: mixtape.PublishedAt.Time,
		SongCount:   mixtape.SongCount,
		AuthorName:  mixtape.AuthorName,
	}
}

func (a *API) GetPublicFeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts := a.store.GetAggregatedPublicPosts()

		songIds := make([]int32, 0)
		for _, post := range posts {
			if post.SongId.Valid {
				songIds = append(songIds, post.SongId.Int32)
			}
		}

		songs := a.store.GetSongsByIdList(songIds, 4)
		songsById := mapSongsById(songs)

		mixtapeIds := make([]int32, 0)
		for _, post := range posts {
			if post.MixtapeId.Valid {
				mixtapeIds = append(mixtapeIds, post.MixtapeId.Int32)
			}
		}

		mixtapes := a.store.GetMixtapePreviewsByIdList(mixtapeIds)
		mixtapesById := mapMixtapesById(mixtapes)

		feedItems := []feedItemJson{}

		for _, post := range posts {
			item := feedItemJson{
				Timestamp: post.Timestamp,
				UserNames: post.UserNames,
			}

			if post.SongId.Valid {
				song := songsById[post.SongId.Int32]
				serialized := serializeSong(song)
				item.Song = &serialized
			}
			if post.MixtapeId.Valid {
				mixtape := mixtapesById[post.MixtapeId.Int32]
				serialized := serializeMixtape(mixtape)
				item.Mixtape = &serialized
			}

			feedItems = append(feedItems, item)
		}

		render.JSON(w, r, feedItems)
	}
}
