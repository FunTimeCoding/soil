package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
	"time"
)

func New(
	p *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered],
	clock func() time.Time,
) *Service {
	return &Service{pool: p, clock: clock}
}
