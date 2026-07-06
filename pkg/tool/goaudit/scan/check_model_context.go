package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
)

func (s *Service) checkModelContext(
	v *virtual_file_system.System,
	path string,
) {
	if !s.ModelContext {
		return
	}

	mc := filepath.Join(path, "model_context")

	if !v.Has(filepath.Join(mc, "mount.go")) {
		s.addConcern(MissingMountKey, MissingMountText, path)
	}

	if !v.Has(filepath.Join(mc, "capture_fail.go")) {
		s.addConcern(MissingCaptureFailKey, MissingCaptureFailText, path)
	}

	if v.Has(filepath.Join(mc, "nested.go")) {
		s.addConcern(StaleNestedKey, StaleNestedText, path)
	}
}
