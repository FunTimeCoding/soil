package pattern

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
)

func New(
	w *workspace.Service,
	c *commit.Service,
) *Service {
	return &Service{workspace: w, commit: c}
}
