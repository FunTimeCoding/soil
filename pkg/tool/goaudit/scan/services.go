package scan

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"sort"
	"strings"
)

func Services(
	v *virtual_file_system.System,
	repo string,
	configuration *Configuration,
) []*Service {
	var result []*Service
	toolDirectory := "pkg/tool"

	for _, name := range v.MustReadDirectory(toolDirectory) {
		if !strings.HasPrefix(name, constant.Go) {
			continue
		}

		if !v.DirectoryExists(filepath.Join(toolDirectory, name)) {
			continue
		}

		s := scanService(
			v,
			filepath.Join(toolDirectory, name),
			name,
			repo,
			configuration,
		)

		if s != nil {
			result = append(result, s)
		}
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result
}
