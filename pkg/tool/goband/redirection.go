package goband

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goband/constant"
	"github.com/funtimecoding/soil/pkg/tool/goband/plan"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/spf13/cobra"
)

func redirection(x *Context) *cobra.Command {
	var instance string
	var confirmed bool
	result := &cobra.Command{
		Use:   "redirection",
		Short: "Switch on the redirection listener and screen service where they are off",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			v := readStatus(x, instance)
			pending := pendingRedirection(v)

			if len(pending) == 0 || !confirmed {
				x.Terminal.EmitValue(
					plan.NewRedirection(v.Instance, pending),
					nil,
				)

				return
			}

			x.Terminal.EmitValue(
				x.Client.EnableRedirection(
					context.Background(),
					client.RedirectionRequest{Instance: &v.Instance},
				),
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
		"execute the writes instead of describing them",
	)

	return result
}
