package service

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/user"
)

func (s *Service) Users() ([]*user.User, error) {
	found, e := s.directory.Search(
		constant.PersonFilter,
		[]string{
			constant.AccountAttribute,
			constant.NameAttribute,
			constant.SurnameAttribute,
			constant.MailAttribute,
			constant.UniqueAttribute,
		},
	)

	if e != nil {
		return nil, e
	}

	var result []*user.User

	for _, record := range found {
		result = append(
			result,
			user.New(
				first(record, constant.AccountAttribute),
				first(record, constant.NameAttribute),
				first(record, constant.SurnameAttribute),
				first(record, constant.MailAttribute),
				first(record, constant.UniqueAttribute),
				record.DistinguishedName,
			),
		)
	}

	return result, nil
}
