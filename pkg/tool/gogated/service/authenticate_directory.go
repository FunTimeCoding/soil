package service

import (
	"github.com/funtimecoding/soil/pkg/directory/types/entry"
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
	found, e := guarded(
		s.directoryGuard,
		func() (*entry.Entry, error) {
			return s.directory.Authenticate(account, password)
		},
	)

	if e != nil {
		if validation.Is(e) || not_found.Is(e) {
			return nil, nil
		}

		return nil, e
	}

	member, e := s.inGroup(found.Account)

	if e != nil {
		return nil, e
	}

	if !member {
		return nil, nil
	}

	row, e := s.store.UserByDirectory(found.Unique)

	if e != nil {
		return nil, e
	}

	if row != nil {
		return row, s.store.RefreshDirectoryUser(
			row,
			found.Account,
			found.Mail,
			found.Name,
		)
	}

	existing, e := s.store.UserByMail(found.Mail)

	if e != nil {
		return nil, e
	}

	if existing != nil {
		return nil, conflict.Exists(constant.SourceLocal, found.Mail)
	}

	return s.store.CreateDirectoryUser(
		found.Unique,
		found.Account,
		found.Mail,
		found.Name,
	)
}
