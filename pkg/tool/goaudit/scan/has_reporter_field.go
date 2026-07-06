package scan

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"strings"
)

func hasReporterField(
	v *virtual_file_system.System,
	serverPath string,
) bool {
	for _, name := range v.MustReadDirectory(serverPath) {
		if !strings.HasSuffix(name, constant.GoExtension) {
			continue
		}

		content := v.ReadString(filepath.Join(serverPath, name))

		if strings.Contains(content, "reporter") {
			return true
		}
	}

	return false
}
