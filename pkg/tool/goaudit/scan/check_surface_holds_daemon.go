package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
	"strings"
)

func (s *Service) checkSurfaceHoldsDaemon(
	v *virtual_file_system.System,
	path string,
) {
	suffix := fmt.Sprintf(
		"/%s/%s",
		filepath.ToSlash(path),
		constant.DaemonDirectory,
	)

	for _, surface := range constant.Surfaces {
		for _, d := range importingDirectories(
			v,
			filepath.Join(path, surface),
			func(i string) bool { return strings.HasSuffix(i, suffix) },
		) {
			s.addConcern(
				constant.SurfaceHoldsDaemonKey,
				fmt.Sprintf(constant.SurfaceHoldsDaemonText, d),
				path,
			)
		}
	}
}
