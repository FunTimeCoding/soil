package gopnsense

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/spf13/cobra"
)

func deleteHost(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var apply bool
	result := &cobra.Command{
		Use:   "delete-host <identifier>",
		Short: "Delete a Dnsmasq host entry",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitDone(c.DeleteHost(context.Background(), a[0], apply))
		},
	}
	result.Flags().BoolVar(
		&apply,
		"apply",
		true,
		"reconfigure Dnsmasq after the write",
	)

	return result
}
