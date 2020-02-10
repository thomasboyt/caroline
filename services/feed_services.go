package services

import (
	"time"

	"github.com/thomasboyt/caroline/models"
	re "github.com/thomasboyt/caroline/resources"
	"github.com/thomasboyt/caroline/store"
)

const CURRENT_USER_ID_PLACEHOLDER = 4

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

func getRelationsForPosts(store *store.Store, posts []models.PostWithConnections) (map[int32]models.SongWithMeta, map[int32]models.MixtapePreview) {
	songIds := make([]int32, 0)
	mixtapeIds := make([]int32, 0)
	for _, post := range posts {
		if post.GetSongId().Valid {
			songIds = append(songIds, post.GetSongId().Int32)
		}
		if post.GetMixtapeId().Valid {
			mixtapeIds = append(mixtapeIds, post.GetMixtapeId().Int32)
		}
	}

	songs := store.GetSongsByIdList(songIds, CURRENT_USER_ID_PLACEHOLDER)
	songsById := mapSongsById(songs)

	mixtapes := store.GetMixtapePreviewsByIdList(mixtapeIds)
	mixtapesById := mapMixtapesById(mixtapes)

	return songsById, mixtapesById
}

func getPostRelationsFromMaps(
	post models.AggregatedPost,
	songsById map[int32]models.SongWithMeta,
	mixtapesById map[int32]models.MixtapePreview,
) (*re.SongJson, *re.MixtapePreviewJson) {

	var serializedSong *re.SongJson
	if post.SongId.Valid {
		song := songsById[post.SongId.Int32]
		serialized := serializeSong(song)
		serializedSong = &serialized
	}

	var serializedMixtape *re.MixtapePreviewJson
	if post.MixtapeId.Valid {
		mixtape := mixtapesById[post.MixtapeId.Int32]
		serialized := serializeMixtape(mixtape)
		serializedMixtape = &serialized
	}

	return serializedSong, serializedMixtape
}

func serializeSong(song models.SongWithMeta) re.SongJson {
	songJson := re.SongJson{
		SongWithMeta: song,
	}

	return songJson
}

func serializeMixtape(mixtape models.MixtapePreview) re.MixtapePreviewJson {
	return re.MixtapePreviewJson{
		MixtapePreview: mixtape,
	}
}

func GetPublicFeed(store *store.Store, beforeTimestamp *time.Time, afterTimestamp *time.Time) []re.FeedItemJson {
	posts := store.GetAggregatedPublicPosts(beforeTimestamp, afterTimestamp, 20)

	// cast []posts -> []postsWithConnections interface
	// https://stackoverflow.com/questions/12994679/slice-of-struct-slice-of-interface-it-implements
	postsWithConnections := make([]models.PostWithConnections, len(posts))
	for i, post := range posts {
		postsWithConnections[i] = post
	}

	songsById, mixtapesById := getRelationsForPosts(store, postsWithConnections)
	feedItems := []re.FeedItemJson{}

	for _, post := range posts {
		item := re.FeedItemJson{
			Timestamp: post.Timestamp,
			UserNames: post.UserNames,
		}

		song, mixtape := getPostRelationsFromMaps(post, songsById, mixtapesById)
		item.Song = song
		item.Mixtape = mixtape

		feedItems = append(feedItems, item)
	}

	return feedItems
}

func GetUserPlaylist(store *store.Store, userId int32, beforeTimestamp *time.Time, afterTimestamp *time.Time) []re.PlaylistItemJson {
	posts := store.GetUserPostsByUserId(userId, beforeTimestamp, afterTimestamp, 20)

	// cast []posts -> []postsWithConnections interface
	// https://stackoverflow.com/questions/12994679/slice-of-struct-slice-of-interface-it-implements
	postsWithConnections := make([]models.PostWithConnections, len(posts))
	for i, post := range posts {
		postsWithConnections[i] = post
	}

	songsById, mixtapesById := getRelationsForPosts(store, postsWithConnections)
	playlistItems := []re.PlaylistItemJson{}

	for _, post := range posts {
		item := re.PlaylistItemJson{
			Timestamp: post.Timestamp,
		}

		song, mixtape := getPostRelationsFromMaps(post, songsById, mixtapesById)
		item.Song = song
		item.Mixtape = mixtape

		playlistItems = append(playlistItems, item)
	}

	return playlistItems
}
