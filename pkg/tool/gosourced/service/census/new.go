package census

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"

func New(w *workspace.Service) *Service {
	return &Service{workspace: w}
}
