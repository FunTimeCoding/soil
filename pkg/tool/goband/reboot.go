package goband

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
)

func reboot(c *band.Client) *cobra.Command {
	var confirmed bool
	result := &cobra.Command{
		Use:   "reboot",
		Short: "Power cycle the machine below the operating system",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			state := c.MustPowerState()
			console.Format(
				"Power state: %s\n",
				band.PowerStateName(state.State),
			)

			if !confirmed {
				console.Line(
					"Would request a hard power cycle - the operating system gets no chance to shut down. Pass --yes to execute.",
				)

				return
			}

			errors.PanicOnError(c.RequestPowerState(constant.PowerCycle))
			console.Line("Power cycle requested.")
		},
	}
	result.Flags().BoolVar(
		&confirmed,
		"yes",
		false,
		"execute the power cycle instead of describing it",
	)

	return result
}
