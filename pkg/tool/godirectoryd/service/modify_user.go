package service

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/user"
)

func (s *Service) ModifyUser(
	account string,
	name string,
	surname string,
	mail string,
) (*user.User, error) {
	attributes := map[string][]string{}

	if name != "" {
		attributes[constant.NameAttribute] = []string{name}
	}

	if surname != "" {
		attributes[constant.SurnameAttribute] = []string{surname}
	}

	if mail != "" {
		attributes[constant.MailAttribute] = []string{mail}
	}

	if len(attributes) == 0 {
		return s.User(account)
	}

	if e := s.directory.Modify(s.userName(account), attributes); e != nil {
		return nil, e
	}

	return s.User(account)
}
