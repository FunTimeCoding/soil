package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func (s *Service) DirectoryAllows(identifier string) bool {
	if s.directory == nil {
		return true
	}

	row, e := s.store.UserByIdentifier(identifier)

	if e != nil || row == nil {
		return false
	}

	if row.Source != constant.SourceDirectory {
		return true
	}

	member, e := s.directory.InGroup(row.Account)

	if e != nil {
		return false
	}

	return member
}
