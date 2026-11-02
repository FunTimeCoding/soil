package scan

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"strings"
)

func treeSourceFiles(
	v *virtual_file_system.System,
	directory string,
) []string {
	var result []string
	prefix := fmt.Sprintf("%s%c", directory, filepath.Separator)

	for _, f := range v.Files() {
		if !strings.HasPrefix(f, prefix) ||
			!strings.HasSuffix(f, soil.GoExtension) ||
			strings.HasSuffix(f, soil.TestSuffix) ||
			isTestdataPath(strings.Split(f, string(filepath.Separator))) {
			continue
		}

		result = append(result, f)
	}

	return result
}
