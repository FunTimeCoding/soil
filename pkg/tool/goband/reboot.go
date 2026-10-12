package goband

import (
	"github.com/funtimecoding/soil/pkg/tool/goband/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/spf13/cobra"
)

func reboot(x *Context) *cobra.Command {
	var instance string
	var confirmed bool
	var forced bool
	result := &cobra.Command{
		Use:   "reboot",
		Short: "Ask the operating system to restart, or power cycle with --force",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			description := constant.RebootDescription

			if forced {
				description = constant.RebootForceDescription
			}

			requestPower(
				x,
				instance,
				client.PowerRequestOperationReboot,
				forced,
				confirmed,
				description,
			)

			if confirmed && !forced {
				x.Terminal.Line(constant.GracefulFollowUp)
			}
		},
	}
	f := result.Flags()
	f.StringVar(
		&instance,
		constant.InstanceArgument,
		"",
		constant.InstanceUsage,
	)
	f.BoolVar(
		&confirmed,
		constant.YesArgument,
		false,
		"execute the restart instead of describing it",
	)
	f.BoolVar(
		&forced,
		constant.ForceArgument,
		false,
		"power cycle below the operating system instead of asking it",
	)

	return result
}
