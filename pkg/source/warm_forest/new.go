package warm_forest

import "github.com/funtimecoding/soil/pkg/source/warm_tree"

func New() *Forest {
	return &Forest{trees: make(map[string]*warm_tree.Tree)}
}
