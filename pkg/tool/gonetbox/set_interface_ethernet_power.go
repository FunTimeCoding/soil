package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func setInterfaceEthernetPower(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "set-interface-ethernet-power [device] [interface] [mode] [type]",
		Short: "Set a device interface's ethernet power: mode pse (supplies) or pd (powered), type such as type2-ieee802.3at; an empty mode clears both",
		Args:  cobra.RangeArgs(3, 4),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			kind := ""

			if len(arguments) == 4 {
				kind = arguments[3]
			}

			t.EmitValue(
				c.SetInterfaceEthernetPower(
					context.Background(),
					arguments[0],
					arguments[1],
					arguments[2],
					kind,
				),
			)
		},
	}
}
