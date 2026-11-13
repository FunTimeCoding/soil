package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"path"
)

func (w *Workspace) marksStore(name string) string {
	return path.Join(
		constant.IndexBuildDirectory,
		buildName(w.build),
		constant.IndexMarksKind,
		name,
	)
}
