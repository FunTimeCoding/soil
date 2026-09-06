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

	if !v.Has(filepath.Join(path, constant.MountFileName)) {
		s.addConcern(
			constant.MissingServiceMountKey,
			constant.MissingServiceMountText,
			path,
		)
	}
}
