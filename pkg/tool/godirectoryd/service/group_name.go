package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
)

func (s *Service) groupName(name string) string {
	return fmt.Sprintf(
		"%s=%s,%s,%s",
		constant.NameAttribute,
		name,
		constant.GroupContainer,
		s.directory.Base(),
	)
}
