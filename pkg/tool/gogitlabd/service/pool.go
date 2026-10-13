package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
)

func (s *Service) Pool() *pool.Pool[instance.Instance, *instance.Instance, face.Forge] {
	return s.pool
}
