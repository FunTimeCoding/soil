package scan

import "github.com/funtimecoding/soil/pkg/tool/goaudit/constant"

func (s *Service) checkConstantFile(path string) {
	if s.ConstantFile {
		s.addConcern(
			constant.ConstantFileKey,
			constant.ConstantFileText,
			path,
		)
	}
}
