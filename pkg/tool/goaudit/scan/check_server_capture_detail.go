package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func (s *Service) checkServerCaptureDetail(
	v *virtual_file_system.System,
	path string,
) {
	if v.Has(filepath.Join(path, "server", "capture_detail.go")) {
		return
	}

	s.addConcern(
		constant.MissingCaptureDetailKey,
		constant.MissingCaptureDetailText,
		path,
	)
}
