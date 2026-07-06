package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
)

func (s *Service) checkStaleDirectories(
	v *virtual_file_system.System,
	path string,
) {
	if v.DirectoryExists(filepath.Join(path, "route")) {
		s.addConcern(StaleRouteKey, StaleRouteText, path)
	}

	if v.DirectoryExists(filepath.Join(path, "tool")) {
		s.addConcern(StaleToolKey, StaleToolText, path)
	}

	if v.DirectoryExists(filepath.Join(path, "toolset")) {
		s.addConcern(StaleToolsetKey, StaleToolsetText, path)
	}

	if v.DirectoryExists(filepath.Join(path, "poller")) {
		s.addConcern(StalePollerKey, StalePollerText, path)
	}
}
