package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"strings"
)

func (s *Service) checkServerCaptureFail(
	v *virtual_file_system.System,
	path string,
) {
	specPath := filepath.Join(path, "generated", "server", "openapi.yaml")

	if !v.Has(specPath) {
		return
	}

	content := v.ReadString(specPath)

	if !strings.Contains(content, "\"500\"") {
		return
	}

	captureFailPath := filepath.Join(path, "server", "capture_fail.go")

	if v.Has(captureFailPath) {
		return
	}

	if !hasReporterField(v, filepath.Join(path, "server")) {
		return
	}

	s.addConcern(
		MissingServerCaptureFailKey,
		MissingServerCaptureFailText,
		path,
	)
}
