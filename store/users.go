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

func (s *Store) GetUserByAuthToken(authToken string) *models.User {
	user := models.User{}

	err := s.db.Get(&user, `
		select users.*
		from users
		join auth_tokens
		on auth_tokens.user_id = users.id
		where auth_token = $1
	`, authToken)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Panic(err)
	}

	return &user
}
