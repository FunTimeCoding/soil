package reference

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
)

type Service struct {
	workspace *workspace.Service
	commit    *commit.Service
}
