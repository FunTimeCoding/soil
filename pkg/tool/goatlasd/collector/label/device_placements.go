package label

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"time"
)

func devicePlacements(
	devices []*client.Device,
	seenAt time.Time,
) []*placement.Placement {
	var result []*placement.Placement

	for _, d := range devices {
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

	return result
}
