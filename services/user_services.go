package services

import (
	r "github.com/thomasboyt/jam-buds-golang/resources"
	"github.com/thomasboyt/jam-buds-golang/store"
)

func GetUserProfileByUserName(s *store.Store, userName string) *r.UserProfileJson {
	user := s.GetUserByUserName(userName)

	if user == nil {
		return nil
	}

	// colorScheme := s.GetColorSchemeByUserId(user.id)

	return &r.UserProfileJson{
		Id:   user.Id,
		Name: user.Name,
		// ColorScheme: colorScheme,
	}
}
