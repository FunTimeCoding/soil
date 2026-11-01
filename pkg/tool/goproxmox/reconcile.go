package goproxmox

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	goproxmox "github.com/funtimecoding/soil/pkg/tool/goproxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
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
			c.Terminal().Emit(c.Client().Reconcile(full, write))
		},
	}
	result.Flags().BoolVar(
		&full,
		constant.FullParameter,
		false,
		"also record stopped guests as offline",
	)
	result.Flags().BoolVar(
		&write,
		goproxmox.YesArgument,
		false,
		"write the changes",
	)

	return result
}
