package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func New(
	p *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered],
) *Service {
	return &Service{pool: p}
}
