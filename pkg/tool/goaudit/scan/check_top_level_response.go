package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func (s *Service) checkTopLevelResponse(
	v *virtual_file_system.System,
	path string,
) {
	if !v.DirectoryExists(filepath.Join(path, "response")) {
		return
	}

	if !s.ModelContext && !s.Server {
		s.addConcern(
			constant.TopLevelResponseOrphanKey,
			constant.TopLevelResponseOrphanText,
			path,
		)
	}

	if s.ModelContext && !s.Convert &&
		!v.DirectoryExists(filepath.Join(path, "server")) {
		s.addConcern(
			constant.TopLevelResponseMCPKey,
			constant.TopLevelResponseMCPText,
			path,
		)
	}
}
