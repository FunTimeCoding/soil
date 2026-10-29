package label

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"time"
)

func virtualPlacements(
	machines []*client.VirtualMachine,
	seenAt time.Time,
) []*placement.Placement {
	var result []*placement.Placement

	for _, m := range machines {
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

	return result
}
