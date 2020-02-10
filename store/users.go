package store

import (
	"database/sql"
	"log"

	"github.com/thomasboyt/caroline/models"
)

func (s *Store) GetUserByUserName(userName string) *models.User {
	user := models.User{}

	args := map[string]interface{}{
		"userName": userName,
	}

	stmt, err := s.db.PrepareNamed(`
		select users.*
		from users
		where name = :userName;
	`)

	err = stmt.Get(&user, args)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Panic(err)
	}

	return &user
}
