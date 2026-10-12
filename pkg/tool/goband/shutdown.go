package goband

import (
	"github.com/funtimecoding/soil/pkg/tool/goband/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/spf13/cobra"
)

func shutdown(x *Context) *cobra.Command {
	var instance string
	var confirmed bool
	var forced bool
	result := &cobra.Command{
		Use:   "shutdown",
		Short: "Ask the operating system to shut down, or cut power with --force",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			description := constant.ShutdownDescription

			if forced {
				description = constant.ShutdownForceDescription
			}

			requestPower(
				x,
				instance,
				client.PowerRequestOperationShutdown,
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
		"execute the shutdown instead of describing it",
	)
	f.BoolVar(
		&forced,
		constant.ForceArgument,
		false,
		"cut power below the operating system instead of asking it",
	)

	return result
}
