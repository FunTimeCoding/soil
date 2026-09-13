package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/user"
)

func (s *Service) User(account string) (*user.User, error) {
	found, e := s.Users()

	if e != nil {
		return nil, e
	}

	for _, u := range found {
		if u.Account == account {
			return u, nil
		}
	}

	return nil, not_found.New(constant.UserKind, account)
}
