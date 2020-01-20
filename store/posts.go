package store

import (
	"log"

	"github.com/thomasboyt/jam-buds-golang/models"
)

// TODO: Limit + offset + etc.
func (s *Store) GetAggregatedPublicPosts() []models.AggregatedPost {
	posts := []models.AggregatedPost{}

	err := s.db.Select(&posts, `
		SELECT
			song_id,
			mixtape_id,
			MIN(posts.created_at) as timestamp,
			ARRAY_AGG(users.name) as user_names
		FROM
			posts
		JOIN
			users ON users.id = posts.user_id
		WHERE
			users.show_in_public_feed = true
		GROUP BY
			posts.song_id, posts.mixtape_id
		ORDER BY
			timestamp DESC;
	`)

	if err != nil {
		log.Panic(err)
	}

	return posts
}
