package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func (s *Service) checkServiceMount(
	v *virtual_file_system.System,
	path string,
) {
	if !s.guarded() {
		return
	}

	owner := path

	if daemon := filepath.Join(path, constant.DaemonDirectory); v.DirectoryExists(
		daemon,
	) {
		owner = daemon
	}

	if !v.Has(filepath.Join(owner, constant.MountFileName)) {
		s.addConcern(
			constant.MissingServiceMountKey,
			constant.MissingServiceMountText,
			path,
		)
	}
}
