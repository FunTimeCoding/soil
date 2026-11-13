package option

import (
	"github.com/funtimecoding/soil/pkg/lint/document_cache"
	"github.com/funtimecoding/soil/pkg/lint/pointer_cache"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"io"
)

type Lint struct {
	Layer        *layer.Layer
	Trees        *warm_forest.Forest
	Documents    *document_cache.Cache
	Resolved     *pointer_cache.Cache
	Trace        io.Writer
	Skips        []string
	ReflowSkips  []string
	Scopes       []string
	Registries   []string
	Verbose      bool
	Census       bool
	Fix          bool
	Summary      bool
	Metadata     bool
	BeforeCommit func()
}
