package goband

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
)

func start(c *band.Client) *cobra.Command {
	var confirmed bool
	result := &cobra.Command{
		Use:   "start",
		Short: "Power on a machine that is off",
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
					"Would power the machine on - the firmware only offers this while it is off. Pass --yes to execute.",
				)

				return
			}

			errors.PanicOnError(c.RequestPowerState(constant.PowerOn))
			console.Line(
				"Power on requested. Poll status until the state reads on.",
			)
		},
	}
	result.Flags().BoolVar(
		&confirmed,
		"yes",
		false,
		"execute the power on instead of describing it",
	)

	return result
}
