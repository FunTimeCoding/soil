package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func (s *Service) Snippets() *pool.Pool[instance.Instance, *instance.Instance, face.SnippetClient] {
	return s.snippets
}
