package gocheck

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maps"
	"slices"
)

func directoryPatterns(directories map[string]bool) []string {
	var result []string

	for _, d := range slices.Sorted(maps.Keys(directories)) {
		if d == soil.CurrentDirectory {
			result = append(result, d)

			continue
		}

		result = append(
			result,
			join.Empty(soil.CurrentDirectory, stringsConstant.Slash, d),
		)
	}

	return result
}
