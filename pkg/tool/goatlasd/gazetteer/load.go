package gazetteer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func Load(
	q context.Context,
	c Source,
) (*Gazetteer, error) {
	devices, e := c.ListDevicesWithResponse(q, &client.ListDevicesParams{})

	if e != nil {
		return nil, e
	}

	if devices.JSON200 == nil {
		return nil, unexpected.Format(
			constant.DeviceStatusFormat,
			devices.HTTPResponse.StatusCode,
		)
	}

	machines, f := c.ListVirtualMachinesWithResponse(q)

	if f != nil {
		return nil, f
	}

	if machines.JSON200 == nil {
		return nil, unexpected.Format(
			constant.VirtualMachineStatusFormat,
			machines.HTTPResponse.StatusCode,
		)
	}

	physical, g := c.ListPhysicalAddressesWithResponse(q)

	if g != nil {
		return nil, g
	}

	if physical.JSON200 == nil {
		return nil, unexpected.Format(
			constant.PhysicalStatusFormat,
			physical.HTTPResponse.StatusCode,
		)
	}

	return New(*devices.JSON200, *machines.JSON200, *physical.JSON200), nil
}
