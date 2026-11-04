package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func mountFiles(path string) []string {
	return []string{
		filepath.Join(path, constant.MountFileName),
		filepath.Join(path, constant.DaemonDirectory, constant.MountFileName),
		filepath.Join(path, "run.go"),
	}
}
