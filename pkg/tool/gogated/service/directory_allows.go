package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func (s *Service) DirectoryAllows(identifier string) (bool, error) {
	if s.directory == nil {
		return true, nil
	}

	row, e := s.store.UserByIdentifier(identifier)

	if e != nil {
		return false, e
	}

	if row == nil {
		return false, nil
	}

	if row.Source != constant.SourceDirectory {
		return true, nil
	}

	return s.inGroup(row.Account)
}
