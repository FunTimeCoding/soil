package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/option"
)

func New(
	o *option.Prometheus,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	p := pool.New(o.Inventory, peeredBuilder(o, l, r))

	for _, i := range o.Inventory.Instances {
		_, e := p.Client(i.Name)
		errors.PanicOnError(e)
	}

	return &Service{
		inventory: o.Inventory,
		pool:      p,
		selection: selection.New(),
	}
}
