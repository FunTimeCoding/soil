package workspace

import "github.com/funtimecoding/soil/pkg/source/types/layer_change"

func (w *Workspace) Changes() []*layer_change.Change {
	return w.layer.Changes()
}
