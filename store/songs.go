package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/thomasboyt/caroline/models"
)

// TODO: Limit + offset + etc.
func (s *Store) GetSongsByIdList(songIds []int32, currentUserId int32) []models.SongWithMeta {
	songs := []models.SongWithMeta{}

	arg := map[string]interface{}{
		"songIds":       songIds,
		"currentUserId": currentUserId,
	}

	query, args, err := sqlx.Named(`
		SELECT
			*,
			(SELECT COUNT(*) FROM likes WHERE likes.song_id=songs.id) AS like_count,
			EXISTS(SELECT 1 FROM likes WHERE likes.user_id=:currentUserId AND likes.song_id=songs.id) AS is_liked
		FROM
			songs
		WHERE
			songs.id IN (:songIds);
	`, arg)

	query, args, err = sqlx.In(query, args...)
	query = s.db.Rebind(query)

	err = s.db.Select(&songs, query, args...)

	if err != nil {
		log.Panic(err)
	}

	return songs
}
