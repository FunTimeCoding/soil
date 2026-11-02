package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func (s *Service) checkMissingDaemon(
	v *virtual_file_system.System,
	path string,
) {
	f := parseWebFile(v, filepath.Join(path, "run.go"))

	if f == nil || !callsPackage(f, constant.LifecyclePackage) {
		return
	}

	s.addConcern(constant.MissingDaemonKey, constant.MissingDaemonText, path)
}
