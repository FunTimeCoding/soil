package scan

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
	"path/filepath"
	"sort"
	"strings"
)

func Services(
	v *virtual_file_system.System,
	repo string,
	configuration *audit_configuration.Configuration,
) []*Service {
	var result []*Service
	toolDirectory := "pkg/tool"

	if !v.DirectoryExists(toolDirectory) {
		return nil
	}

	for _, name := range v.MustReadDirectory(toolDirectory) {
		if !strings.HasPrefix(name, soil.Go) {
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
