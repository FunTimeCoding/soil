package check

import (
	"github.com/funtimecoding/soil/pkg/lint/document_cache"
	"github.com/funtimecoding/soil/pkg/lint/pointer_cache"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"time"
)

type Service struct {
	inventory      *inventory.Inventory
	commit         *commit.Service
	workspace      *workspace.Service
	indexDirectory string
	turns          *turn.Scheduler
	trees          *warm_forest.Forest
	documents      *document_cache.Cache
	resolved       *pointer_cache.Cache
	source         string
	started        time.Time
}
