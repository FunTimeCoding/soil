package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/check"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
)

type Server struct {
	census    *census.Service
	check     *check.Service
	diff      *diff.Service
	module    *module.Service
	workspace *workspace.Service
	reporter  face.Reporter
}
