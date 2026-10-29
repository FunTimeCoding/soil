package service

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
)

func (s *Service) wireless() (string, error) {
	return s.command(
		constant.Sudo,
		constant.SudoNonInteract,
		systemConstant.Wdutil,
		systemConstant.WdutilInformation,
	)
}
