package gazetteer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

func Load(
	q context.Context,
	c face.Client,
) (*Gazetteer, error) {
	devices, e := c.ListDevices(q, nil)

	if e != nil {
		return nil, e
	}

	machines, f := c.ListVirtualMachines(q)

	if f != nil {
		return nil, f
	}

	physical, g := c.ListPhysicalAddresses(q)

	if g != nil {
		return nil, g
	}

	return New(devices, machines, physical), nil
}
