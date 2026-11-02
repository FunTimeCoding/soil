package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/constant"
	goproxmoxdConstant "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func reconcile(c *command_context.Context) *cobra.Command {
	var full bool
	var write bool
	result := &cobra.Command{
		Use:   "reconcile",
		Short: "Compare guests with the power intent netbox records",
		Long:  "Lists the guests whose netbox status should change: running guests recorded offline become active, and with --full, stopped guests recorded active become offline. Guests netbox does not know are listed, never created. Writes only with --yes.",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Terminal().EmitValue(
				c.Client().Reconcile(
					context.Background(),
					&client.ReconcileParams{
						Instance: c.Instance(),
						Full:     &full,
						Write:    &write,
					},
				),
			)
		},
	}
	result.Flags().BoolVar(
		&full,
		goproxmoxdConstant.FullParameter,
		false,
		"also record stopped guests as offline",
	)
	result.Flags().BoolVar(
		&write,
		constant.YesArgument,
		false,
		"write the changes",
	)

	return result
}
