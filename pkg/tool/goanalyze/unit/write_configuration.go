package unit

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	"path/filepath"
)

func writeConfiguration(
	root string,
	content string,
) {
	path := filepath.Join(root, constant.ConfigurationPath)
	system.EnsurePathExists(filepath.Dir(path))
	system.SaveFile(path, content)
}
