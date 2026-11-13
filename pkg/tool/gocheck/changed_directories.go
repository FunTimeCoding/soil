package gocheck

import (
	"github.com/funtimecoding/soil/pkg/source/types/layer_change"
	"path/filepath"
)

func changedDirectories(
	root string,
	changes []*layer_change.Change,
) map[string]bool {
	result := make(map[string]bool)

	for _, c := range changes {
		result[relativeDirectory(root, filepath.Dir(c.Path))] = true
		result[relativeDirectory(root, filepath.Dir(c.Destination))] = true
	}

	return result
}
