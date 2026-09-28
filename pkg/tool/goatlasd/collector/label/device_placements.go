package label

import (
	"context"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"time"
)

func (c *Collector) devicePlacements(
	q context.Context,
	seenAt time.Time,
) ([]*placement.Placement, error) {
	var result []*placement.Placement
	devices, e := c.netbox.ListDevicesWithResponse(
		q,
		&client.ListDevicesParams{},
	)

	if e != nil {
		return nil, e
	}

	if devices.JSON200 == nil {
		return nil, statusFail(
			constant.DeviceStatusFormat,
			devices.HTTPResponse.StatusCode,
		)
	}

	for _, d := range *devices.JSON200 {
		for _, name := range ServiceNames(d.Labels) {
			result = append(
				result,
				placement.New(
					constant.SourceNetbox,
					constant.KindService,
					"",
					name,
					place.New(
						netboxConstant.DeviceAddress,
						d.Identifier,
						d.Name,
					),
					seenAt,
				),
			)
		}
	}

	return result, nil
}
