package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func deleteHost(c *client.Client) *cobra.Command {
	var apply bool
	result := &cobra.Command{
		Use:   "delete-host <identifier>",
		Short: "Delete a Dnsmasq host entry",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			console.Emit(c.DeleteHost(a[0], &apply))
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
