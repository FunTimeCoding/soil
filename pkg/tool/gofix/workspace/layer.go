package workspace

import "github.com/funtimecoding/soil/pkg/source/layer"

func (w *Workspace) Layer() *layer.Layer {
	return w.layer
}
