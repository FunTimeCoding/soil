package service

import "github.com/funtimecoding/soil/pkg/chat/teams/user"

func (s *Service) UserProfile(input string) (*user.User, error) {
	return s.resolveUser(input)
}
