package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
)

func New(
	p *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered],
) *Service {
	return &Service{pool: p}
}
