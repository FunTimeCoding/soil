package scan

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"sort"
	"strings"
)

func MisplacedTests(v *virtual_file_system.System) []*concern.Concern {
	var result []*concern.Concern

	for _, path := range v.Files() {
		if !strings.HasSuffix(path, constant.TestSuffix) {
			continue
		}

		if isTestHomePath(path) {
			continue
		}

		result = append(
			result,
			concern.NewFile(
				MisplacedTestKey,
				MisplacedTestText,
				path,
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
