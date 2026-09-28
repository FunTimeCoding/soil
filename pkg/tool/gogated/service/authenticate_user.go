package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/user"
)

func (s *Service) AuthenticateUser(
	mail string,
	password string,
	source string,
) (*user.User, error) {
	if s.directory == nil || source != constant.SourceDirectory {
		return s.store.AuthenticateUser(mail, password)
	}

	return s.authenticateDirectory(mail, password)
}
