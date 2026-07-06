package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"sort"
)

func MissingSentry(v *virtual_file_system.System) []*concern.Concern {
	programs := make(map[string]bool)

	for _, name := range v.MustReadDirectory("cmd") {
		if name == "example" {
			continue
		}

		if !v.DirectoryExists(filepath.Join("cmd", name)) {
			continue
		}

		programs[name] = false
	}

	for name := range programs {
		toolDirectory := filepath.Join("pkg", "tool", name)

		if hasSentryReporter(v, toolDirectory) {
			programs[name] = true
		}
	}

	var result []*concern.Concern

	for name, hasSentry := range programs {
		if !hasSentry {
			result = append(
				result,
				concern.NewPackage(
					MissingSentryKey,
					MissingSentryText,
					fmt.Sprintf("cmd/%s", name),
				),
			)
		}
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Path < result[j].Path
		},
	)

	return result
}
