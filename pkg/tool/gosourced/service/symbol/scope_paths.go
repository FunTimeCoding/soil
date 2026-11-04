package symbol

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"sort"
	"strings"
)

func (s *Service) scopePaths(
	directory string,
	packagePath string,
) []string {
	prefix, subtree := strings.CutSuffix(packagePath, "/...")

	if !subtree {
		return []string{packagePath}
	}

	var result []string

	for path := range s.workspace.Workspace(directory).Graph().Units {
		if path == prefix || strings.HasPrefix(path, join.Empty(prefix, "/")) {
			result = append(result, path)
		}
	}

	sort.Strings(result)

	return result
}
