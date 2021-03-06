package store

import (
	"log"
	"time"

	"github.com/thomasboyt/caroline/models"
)

func (s *Store) GetAggregatedPublicPosts(
	beforeTimestamp *time.Time,
	afterTimestamp *time.Time,
	limit int32) []models.AggregatedPost {

	posts := []models.AggregatedPost{}

	arg := map[string]interface{}{
		"beforeTimestamp": beforeTimestamp,
		"afterTimestamp":  afterTimestamp,
		"limit":           limit,
	}

	stmt, err := s.db.PrepareNamed(`
		SELECT
			song_id,
			mixtape_id,
			MIN(posts.created_at) as "timestamp",
			ARRAY_AGG(users.name) as user_names
		FROM posts
		JOIN users ON users.id = posts.user_id
		WHERE
			users.show_in_public_feed = true
		GROUP BY posts.song_id, posts.mixtape_id
		HAVING
			(CAST(:beforeTimestamp AS timestamp) IS NULL OR MIN(posts.created_at) < :beforeTimestamp)
			AND
			(CAST(:afterTimestamp AS timestamp) IS NULL OR MIN(posts.created_at) > :afterTimestamp)
		ORDER BY timestamp DESC
		LIMIT
			CASE
				WHEN CAST(:afterTimestamp AS TIMESTAMP) IS NULL
				THEN CAST(:limit AS bigint)
				ELSE NULL
			END;
	`)

	if err != nil {
		log.Panic(err)
	}

	err = stmt.Select(&posts, arg)

	if err != nil {
		log.Panic(err)
	}

	return posts
}

func (s *Store) GetUserPostsByUserId(
	userId int32,
	beforeTimestamp *time.Time,
	afterTimestamp *time.Time,
	limit int32) []models.AggregatedPost {

	posts := []models.AggregatedPost{}

	arg := map[string]interface{}{
		"userId":          userId,
		"beforeTimestamp": beforeTimestamp,
		"afterTimestamp":  afterTimestamp,
		"limit":           limit,
		"onlyMixtapes":    false, // TODO
	}

	stmt, err := s.db.PrepareNamed(`
		SELECT
			posts.song_id,
			posts.mixtape_id,
			posts.created_at as "timestamp"
		FROM posts
		JOIN users ON users.id = posts.user_id
		WHERE
			posts.user_id = :userId
			AND (CAST(:beforeTimestamp AS timestamp) IS NULL OR posts.created_at < :beforeTimestamp)
			AND (CAST(:afterTimestamp AS timestamp) IS NULL OR posts.created_at > :afterTimestamp)
			AND (NOT :onlyMixtapes OR posts.mixtape_id IS NOT NULL)
		ORDER BY timestamp DESC
		LIMIT
			CASE
				WHEN CAST(:afterTimestamp AS TIMESTAMP) IS NULL
				THEN CAST(:limit AS bigint)
				ELSE NULL
			END;
	`)

	if err != nil {
		log.Panic(err)
	}

	err = stmt.Select(&posts, arg)

	if err != nil {
		log.Panic(err)
	}

	return posts
}
