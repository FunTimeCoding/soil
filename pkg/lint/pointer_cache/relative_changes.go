package pointer_cache

import (
	"path"
	"path/filepath"
)

func relativeChanges(
	root string,
	changes map[string][]string,
) []string {
	var result []string

	for tree, paths := range changes {
		relative, e := filepath.Rel(root, tree)

		if e != nil {
			continue
		}

		relative = filepath.ToSlash(relative)

		for _, p := range paths {
			result = append(result, path.Join(relative, p))
		}
	}

	return result
}
