package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"path/filepath"
)

func settingPath(root string) string {
	return filepath.Join(
		root,
		constant.SettingsDirectory,
		constant.SettingsFile,
	)
}
