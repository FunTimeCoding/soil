package warm_forest

import (
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"path/filepath"
)

func (f *Forest) Tree(root string) *warm_tree.Tree {
	clean := filepath.Clean(root)

	if t, found := f.trees[clean]; found {
		return t
	}

	result := warm_tree.New(clean)
	f.trees[clean] = result

	return result
}
