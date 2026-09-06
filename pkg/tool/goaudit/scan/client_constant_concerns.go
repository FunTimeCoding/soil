package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"maps"
	"path/filepath"
	"slices"
	"strings"
)

func clientConstantConcerns(
	v *virtual_file_system.System,
	directory string,
) []*concern.Concern {
	values := make(map[string]bool)

	for _, name := range v.MustReadDirectory(directory) {
		f := parseWebFile(v, filepath.Join(directory, name))

		if f == nil {
			continue
		}

		for _, value := range stringConstants(f) {
			values[value] = true
		}
	}

	var result []*concern.Concern

	for _, value := range slices.Sorted(maps.Keys(values)) {
		prefix, found := strings.CutSuffix(value, "_TOKEN")

		if !found || !strings.HasPrefix(prefix, "GO") {
			continue
		}

		for _, suffix := range []string{"_HOST", "_PORT", "_INSECURE"} {
			if !values[fmt.Sprintf("%s%s", prefix, suffix)] {
				result = append(
					result,
					concern.NewPackage(
						constant.ClientConstantKey,
						fmt.Sprintf(
							"%s has no %s%s sibling in the same constant package",
							value,
							prefix,
							suffix,
						),
						directory,
					),
				)
			}
		}
	}

	return result
}
