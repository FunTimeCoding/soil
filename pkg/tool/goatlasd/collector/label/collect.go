package label

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"time"
)

func (c *Collector) Collect(
	q context.Context,
	_ *gazetteer.Gazetteer,
) ([]*placement.Placement, error) {
	var devices []*client.Device
	var machines []*client.VirtualMachine
	e := c.guard.Attempt(
		q,
		func(x context.Context) error {
			var f error

			if devices, f = c.netbox.ListDevices(x, nil); f != nil {
				return f
			}

			machines, f = c.netbox.ListVirtualMachines(x)

			return f
		},
	)

	if e != nil {
		return nil, e
	}

	now := time.Now()

	return append(
		devicePlacements(devices, now),
		virtualPlacements(machines, now)...,
	), nil
}
