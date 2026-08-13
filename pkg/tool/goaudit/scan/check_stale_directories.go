package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func (s *Service) checkStaleDirectories(
	v *virtual_file_system.System,
	path string,
) {
	if v.DirectoryExists(filepath.Join(path, "route")) {
		s.addConcern(constant.StaleRouteKey, constant.StaleRouteText, path)
	}

	if v.DirectoryExists(filepath.Join(path, constant.ToolDirectory)) {
		s.addConcern(constant.StaleToolKey, constant.StaleToolText, path)
	}

	if v.DirectoryExists(filepath.Join(path, "toolset")) {
		s.addConcern(constant.StaleToolsetKey, constant.StaleToolsetText, path)
	}

	if v.DirectoryExists(filepath.Join(path, "poller")) {
		s.addConcern(constant.StalePollerKey, constant.StalePollerText, path)
	}
}
