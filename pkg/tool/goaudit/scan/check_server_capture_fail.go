package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
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

	server := filepath.Join(path, "server")

	if v.Has(filepath.Join(server, constant.CaptureFile)) ||
		v.Has(filepath.Join(server, constant.AnswerWithFile)) {
		return
	}

	if !hasReporterField(v, server) {
		return
	}

	s.addConcern(
		constant.MissingServerCaptureFailKey,
		constant.MissingServerCaptureFailText,
		path,
	)
}
