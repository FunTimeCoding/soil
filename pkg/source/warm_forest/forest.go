package warm_forest

import "github.com/funtimecoding/soil/pkg/source/warm_tree"

type Forest struct {
	trees map[string]*warm_tree.Tree
}
