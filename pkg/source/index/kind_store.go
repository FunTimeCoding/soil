package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"path"
)

func (w *Workspace) kindStore(
	base string,
	k *kind.Kind,
) string {
	return path.Join(
		constant.IndexBuildDirectory,
		buildName(w.build),
		base,
		k.Name,
	)
}
