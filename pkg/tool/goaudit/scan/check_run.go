package scan

import "github.com/funtimecoding/soil/pkg/tool/goaudit/constant"

func (s *Service) checkRun(path string) {
	if !s.Run {
		s.addConcern(constant.MissingRunKey, constant.MissingRunText, path)
	}
}
