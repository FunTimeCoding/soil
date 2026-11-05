package goband

import (
	"github.com/funtimecoding/soil/pkg/tool/goband/constant"
	"github.com/spf13/cobra"
)

func status(x *Context) *cobra.Command {
	var instance string
	result := &cobra.Command{
		Use:   "status",
		Short: "Show provisioning, redirection, consent and power state",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			x.Terminal.EmitValue(readStatus(x, instance), nil)
		},
	}
	result.Flags().StringVar(
		&instance,
		constant.InstanceArgument,
		"",
		constant.InstanceUsage,
	)

	return result
}
