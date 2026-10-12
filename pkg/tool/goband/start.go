package goband

import (
	"github.com/funtimecoding/soil/pkg/tool/goband/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/spf13/cobra"
)

func start(x *Context) *cobra.Command {
	var instance string
	var confirmed bool
	result := &cobra.Command{
		Use:   "start",
		Short: "Power on a machine that is off",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			requestPower(
				x,
				instance,
				client.PowerRequestOperationStart,
				false,
				confirmed,
				constant.StartDescription,
			)
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
		"execute the power on instead of describing it",
	)

	return result
}
