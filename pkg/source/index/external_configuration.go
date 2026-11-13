package index

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func (w *Workspace) externalConfiguration(
	mode packages.LoadMode,
) *packages.Config {
	return &packages.Config{
		Mode:       mode,
		Dir:        w.root,
		BuildFlags: resolve.TagFlags(w.graph.Tags),
	}
}
