package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/thomasboyt/caroline/models"
)

func (s *Store) GetMixtapePreviewsByIdList(mixtapeIds []int32) []models.MixtapePreview {
	mixtapes := []models.MixtapePreview{}

	arg := map[string]interface{}{
		"mixtapeIds": mixtapeIds,
	}

	query, args, err := sqlx.Named(`
		SELECT
			mixtapes.*,
			(SELECT users.name FROM users WHERE users.id=mixtapes.user_id) as author_name,
			(SELECT COUNT (*) FROM mixtape_song_entries WHERE mixtape_id=mixtapes.id) as song_count
		FROM
			mixtapes
		JOIN
			users ON users.id = mixtapes.user_id
		WHERE
			mixtapes.id IN (:mixtapeIds);
	`, arg)

	query, args, err = sqlx.In(query, args...)
	query = s.db.Rebind(query)

	err = s.db.Select(&mixtapes, query, args...)

	if err != nil {
		log.Panic(err)
	}

	return mixtapes
}
