package service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/user"
)

func (s *Service) authenticateDirectory(
	account string,
	password string,
) (*user.User, error) {
	entry, e := s.directory.Authenticate(account, password)

	if e != nil {
		if validation.Is(e) || not_found.Is(e) {
			return nil, nil
		}

		return nil, e
	}

	member, e := s.directory.InGroup(entry.Account)

	if e != nil {
		return nil, e
	}

	if !member {
		return nil, nil
	}

	row, e := s.store.UserByDirectory(entry.Unique)

	if e != nil {
		return nil, e
	}

	if row != nil {
		return row, s.store.RefreshDirectoryUser(
			row,
			entry.Account,
			entry.Mail,
			entry.Name,
		)
	}

	existing, e := s.store.UserByMail(entry.Mail)

	if e != nil {
		return nil, e
	}

	if existing != nil {
		return nil, conflict.Exists(constant.SourceLocal, entry.Mail)
	}

	return s.store.CreateDirectoryUser(
		entry.Unique,
		entry.Account,
		entry.Mail,
		entry.Name,
	)
}
