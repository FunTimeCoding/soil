package workspace

import "github.com/funtimecoding/soil/pkg/source/layer"

func Over(
	parent *layer.Layer,
	diff bool,
) *Workspace {
	return &Workspace{
		diff:    diff,
		stacked: true,
		roots:   parent.Roots(),
		layer:   layer.New(parent),
	}
}
