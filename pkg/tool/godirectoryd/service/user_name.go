package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
)

func (s *Service) userName(account string) string {
	return fmt.Sprintf(
		"%s=%s,%s,%s",
		constant.AccountAttribute,
		account,
		constant.UserContainer,
		s.directory.Base(),
	)
}
