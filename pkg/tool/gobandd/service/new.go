package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/peered"
)

func New(
	p *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered],
	l *logger.Logger,
	r face.Reporter,
) *Service {
	return &Service{pool: p, logger: l, reporter: r}
}
