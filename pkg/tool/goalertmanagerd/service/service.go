package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

type Service struct {
	pool *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
}
