package service

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/user"
)

func (s *Service) CreateUser(
	account string,
	name string,
	surname string,
	mail string,
	password string,
) (*user.User, error) {
	attributes := map[string][]string{
		constant.ClassAttribute:   {constant.PersonClass},
		constant.AccountAttribute: {account},
		constant.NameAttribute:    {name},
		constant.SurnameAttribute: {surname},
	}

	if mail != "" {
		attributes[constant.MailAttribute] = []string{mail}
	}

	if password != "" {
		attributes[constant.PasswordAttribute] = []string{password}
	}

	if e := s.directory.Add(s.userName(account), attributes); e != nil {
		return nil, e
	}

	return s.User(account)
}
