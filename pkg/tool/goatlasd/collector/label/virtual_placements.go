package label

import (
	"context"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func (c *Collector) virtualPlacements(
	q context.Context,
	seenAt time.Time,
) ([]*placement.Placement, error) {
	var result []*placement.Placement
	machines, e := c.netbox.ListVirtualMachinesWithResponse(q)

	if e != nil {
		return nil, e
	}

	if machines.JSON200 == nil {
		return nil, statusFail(
			constant.VirtualMachineStatusFormat,
			machines.HTTPResponse.StatusCode,
		)
	}

	for _, m := range *machines.JSON200 {
		for _, name := range ServiceNames(m.Labels) {
			result = append(
				result,
				placement.New(
					constant.SourceNetbox,
					constant.KindService,
					"",
					name,
					place.New(
						netboxConstant.VirtualMachineAddress,
						m.Identifier,
						m.Name,
					),
					seenAt,
				),
			)
		}
	}

	return result, nil
}
