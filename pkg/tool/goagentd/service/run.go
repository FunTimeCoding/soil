package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"strings"
)

func (s *Service) Run(intent string) error {
	if strings.TrimSpace(intent) == "" {
		return validation.New(constant.IntentRequired)
	}

	return s.runner.Begin(intent)
}
