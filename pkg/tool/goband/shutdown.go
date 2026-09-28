package goband

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
)

func shutdown(c *band.Client) *cobra.Command {
	var confirmed bool
	result := &cobra.Command{
		Use:   "shutdown",
		Short: "Press the power button and let the operating system answer",
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
					"Would assert the power button - a live operating system shuts down cleanly, a wedged one ignores it. Pass --yes to execute.",
				)

				return
			}

			errors.PanicOnError(c.RequestPowerState(constant.PowerOffSoft))
			console.Line(
				"Power button asserted. Poll status - if the state stays on, the operating system is not answering and reboot is the remaining option.",
			)
		},
	}
	result.Flags().BoolVar(
		&confirmed,
		"yes",
		false,
		"execute the shutdown instead of describing it",
	)

	return result
}
