package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func (s *Service) Pool() *pool.Pool[instance.Instance, *instance.Instance, face.ProxmoxClient] {
	return s.pool
}
