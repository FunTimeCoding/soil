package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
)

func (s *Service) Pool() *pool.Pool[instance.Instance, *instance.Instance, face.Engine] {
	return s.pool
}
