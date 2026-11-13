package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/check"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/file"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/packages"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
	"sync"
)

type Daemon struct {
	memory    face.Client
	guard     *peer.Guard
	census    *census.Service
	check     *check.Service
	diff      *diff.Service
	file      *file.Service
	module    *module.Service
	packages  *packages.Service
	pattern   *pattern.Service
	reference *reference.Service
	symbol    *symbol.Service
	workspace *workspace.Service
	closeOnce sync.Once
}
