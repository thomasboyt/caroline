package services

import (
	"time"

	"github.com/thomasboyt/jam-buds-golang/models"
	r "github.com/thomasboyt/jam-buds-golang/resources"
	"github.com/thomasboyt/jam-buds-golang/store"
)

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

func getPostRelations(
	post models.AggregatedPost,
	songsById map[int32]models.SongWithMeta,
	mixtapesById map[int32]models.MixtapePreview,
) (*r.SongJson, *r.MixtapePreviewJson) {

	var serializedSong *r.SongJson
	if post.SongId.Valid {
		song := songsById[post.SongId.Int32]
		serialized := serializeSong(song)
		serializedSong = &serialized
	}

	var serializedMixtape *r.MixtapePreviewJson
	if post.MixtapeId.Valid {
		mixtape := mixtapesById[post.MixtapeId.Int32]
		serialized := serializeMixtape(mixtape)
		serializedMixtape = &serialized
	}

	return serializedSong, serializedMixtape
}

func serializeSong(song models.SongWithMeta) r.SongJson {
	songJson := r.SongJson{
		SongWithMeta: song,
	}

	return songJson
}

func serializeMixtape(mixtape models.MixtapePreview) r.MixtapePreviewJson {
	return r.MixtapePreviewJson{
		MixtapePreview: mixtape,
	}
}

func GetPublicFeed(store *store.Store, beforeTimestamp *time.Time, afterTimestamp *time.Time) []r.FeedItemJson {
	posts := store.GetAggregatedPublicPosts(beforeTimestamp, afterTimestamp, 20)

	songIds := make([]int32, 0)
	for _, post := range posts {
		if post.SongId.Valid {
			songIds = append(songIds, post.SongId.Int32)
		}
	}

	songs := store.GetSongsByIdList(songIds, 4)
	songsById := mapSongsById(songs)

	mixtapeIds := make([]int32, 0)
	for _, post := range posts {
		if post.MixtapeId.Valid {
			mixtapeIds = append(mixtapeIds, post.MixtapeId.Int32)
		}
	}

	mixtapes := store.GetMixtapePreviewsByIdList(mixtapeIds)
	mixtapesById := mapMixtapesById(mixtapes)

	feedItems := []r.FeedItemJson{}

	for _, post := range posts {
		item := r.FeedItemJson{
			Timestamp: post.Timestamp,
			UserNames: post.UserNames,
		}

		song, mixtape := getPostRelations(post, songsById, mixtapesById)
		item.Song = song
		item.Mixtape = mixtape

		feedItems = append(feedItems, item)
	}

	return feedItems
}
