package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
	"strings"
)

func isExcluded(
	root string,
	path string,
	configuration *audit_configuration.Configuration,
) bool {
	relative := strings.TrimPrefix(path, fmt.Sprintf("%s/", root))

	if strings.HasPrefix(relative, constant.ToolDirectory) {
		return true
	}

	prefixed := fmt.Sprintf("pkg/%s", relative)

	for _, exclude := range configuration.Exclude {
		if strings.HasPrefix(prefixed, exclude) {
			return true
		}
	}

	return false
}
