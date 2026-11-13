package check

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lint/document_cache"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"runtime/debug"
	"time"
)

func New(
	i *inventory.Inventory,
	c *commit.Service,
	w *workspace.Service,
	indexDirectory string,
	r face.Reporter,
) *Service {
	var source string

	if b, okay := debug.ReadBuildInfo(); okay {
		source, _ = i.ModuleDirectory(b.Main.Path)
	}

	return &Service{
		inventory:      i,
		commit:         c,
		workspace:      w,
		indexDirectory: indexDirectory,
		turns:          turn.New(constant.CheckBound),
		trees:          warm_forest.New(),
		documents:      document_cache.New(),
		resolved:       resolvedCache(r),
		source:         source,
		started:        time.Now(),
	}
}
