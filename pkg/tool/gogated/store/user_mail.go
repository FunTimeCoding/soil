package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/user"

func (s *Store) UserMail(identifier string) string {
	var row user.User
	e := s.mapper.Where("identifier = ?", identifier).First(&row).Error

	if e != nil {
		return ""
	}

	return row.Mail
}
