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
		Short: "Set a device interface's ethernet power: mode pse (supplies) or pd (powered), type such as type2-ieee802.3at; no mode clears both",
		Args:  cobra.RangeArgs(2, 4),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			mode := ""

			if len(arguments) > 2 {
				mode = arguments[2]
			}

			kind := ""

			if len(arguments) > 3 {
				kind = arguments[3]
			}

			t.EmitValue(
				c.SetInterfaceEthernetPower(
					context.Background(),
					arguments[0],
					arguments[1],
					mode,
					kind,
				),
			)
		},
	}
}
