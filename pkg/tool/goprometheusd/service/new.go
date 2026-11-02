package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
)

func New(
	p *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered],
) *Service {
	return &Service{pool: p}
}
