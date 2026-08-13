package scan

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
	"sort"
	"strings"
)

func IdentityWarnings(
	v *virtual_file_system.System,
	services []*Service,
) []*concern.Concern {
	checked := map[string]bool{}

	for _, s := range services {
		checked[s.Name] = true
	}

	var result []*concern.Concern
	toolDirectory := "pkg/tool"

	for _, name := range v.MustReadDirectory(toolDirectory) {
		if !strings.HasPrefix(name, library.Go) {
			continue
		}

		if checked[name] {
			continue
		}

		path := filepath.Join(toolDirectory, name)

		if !v.DirectoryExists(filepath.Join(path, constant.ConstantDirectory)) {
			continue
		}

		result = append(result, IdentityConcerns(v, path, name)...)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Path < result[j].Path
		},
	)

	return result
}
