package permission

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"slices"
	"sort"
)

func newestCommands(prompts []*Prompt) []string {
	sorted := slices.Clone(prompts)
	sort.SliceStable(
		sorted,
		func(i, j int) bool { return sorted[i].When.After(sorted[j].When) },
	)
	var result []string

	for _, p := range sorted {
		if p.Command == "" || slices.Contains(result, p.Command) {
			continue
		}

		result = append(result, p.Command)

		if len(result) == constant.PermissionSampleCommands {
			break
		}
	}

	return result
}
