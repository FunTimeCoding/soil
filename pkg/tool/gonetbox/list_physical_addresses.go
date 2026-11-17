package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func listPhysicalAddresses(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-physical-addresses",
		Short: "List every MAC address and the device or virtual machine carrying it",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.ListPhysicalAddresses(context.Background()))
		},
	}
}
