package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func (s *Service) checkTopLevelArgument(
	v *virtual_file_system.System,
	path string,
) {
	if v.DirectoryExists(filepath.Join(path, "argument")) {
		s.addConcern(
			constant.TopLevelArgumentKey,
			constant.TopLevelArgumentText,
			path,
		)
	}
}
