package scan

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path"
	"sort"
	"strings"
)

func StrayHelpers(v *virtual_file_system.System) []*concern.Concern {
	homes := map[string]bool{}

	for _, p := range v.Files() {
		if strings.HasSuffix(p, soil.TestSuffix) && underTestHome(p) {
			homes[path.Dir(p)] = true
		}
	}

	var result []*concern.Concern

	for _, p := range v.Files() {
		if !strings.HasSuffix(p, soil.GoExtension) ||
			strings.HasSuffix(p, soil.TestSuffix) ||
			!homes[path.Dir(p)] {
			continue
		}

		result = append(
			result,
			concern.NewFile(
				constant.StrayHelperKey,
				constant.StrayHelperText,
				p,
				false,
			),
		)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Path < result[j].Path
		},
	)

	return result
}
