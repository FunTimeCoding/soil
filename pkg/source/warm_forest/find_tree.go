package warm_forest

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"strings"
)

func (f *Forest) FindTree(path string) (*warm_tree.Tree, string, bool) {
	clean := filepath.Clean(path)
	var result *warm_tree.Tree

	for root, t := range f.trees {
		inside := strings.HasPrefix(
			clean,
			join.Empty(root, string(filepath.Separator)),
		)

		if (clean == root || inside) &&
			(result == nil || len(root) > len(result.Root())) {
			result = t
		}
	}

	if result == nil {
		return nil, "", false
	}

	if clean == result.Root() {
		return result, soil.CurrentDirectory, true
	}

	relative, e := filepath.Rel(result.Root(), clean)

	if e != nil {
		return nil, "", false
	}

	return result, filepath.ToSlash(relative), true
}
