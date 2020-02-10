package services

import (
	re "github.com/thomasboyt/caroline/resources"
	"github.com/thomasboyt/caroline/store"
)

func GetUserProfileByUserName(s *store.Store, userName string) *re.UserProfileJson {
	user := s.GetUserByUserName(userName)

	if user == nil {
		return nil
	}

	// TODO
	// colorScheme := s.GetColorSchemeByUserId(user.id)

	return &re.UserProfileJson{
		Id:   user.Id,
		Name: user.Name,
		// ColorScheme: colorScheme,
	}
}
