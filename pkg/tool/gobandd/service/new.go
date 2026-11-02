package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func New(
	o *option.Band,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	result := &Service{
		inventory: o.Inventory,
		pool:      pool.New(o.Inventory, newEngine),
		selection: selection.New(),
		logger:    l,
		reporter:  r,
	}

	if o.Netbox == nil {
		return result
	}

	result.netbox = o.Netbox()
	result.netboxGuard = peer.New(
		constant.Identity.Name(),
		result.netbox.Address(),
		result.netbox,
		l,
		r,
	)

	return result
}
