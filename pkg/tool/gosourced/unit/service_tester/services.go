package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/file"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/packages"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
)

type Services struct {
	Workspace *workspace.Service
	Commit    *commit.Service
	Diff      *diff.Service
	File      *file.Service
	Module    *module.Service
	Packages  *packages.Service
	Pattern   *pattern.Service
	Reference *reference.Service
	Symbol    *symbol.Service
}
