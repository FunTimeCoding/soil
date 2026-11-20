package goclaude

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"path/filepath"
	"slices"
)

func LabelAllowance(
	configuration string,
	project string,
) []string {
	paths := slices.Clone(constant.ManagedSettingsFiles)

	for _, f := range constant.SettingsFiles {
		paths = append(paths, filepath.Join(configuration, f))

		if project != "" {
			paths = append(
				paths,
				filepath.Join(project, constant.ClaudeDirectory, f),
			)
		}
	}

	var result []string

	for _, p := range paths {
		if allowsLabel(p) {
			result = append(result, p)
		}
	}

	return result
}
