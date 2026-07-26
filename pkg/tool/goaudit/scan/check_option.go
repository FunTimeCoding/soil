package scan

import "github.com/funtimecoding/soil/pkg/tool/goaudit/constant"

func (s *Service) checkOption(path string) {
	if !s.Option {
		s.addConcern(constant.MissingOptionKey, constant.MissingOptionText, path)
	}
}
