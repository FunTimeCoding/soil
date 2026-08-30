package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func setHost(c *client.Client) *cobra.Command {
	f := &hostFlags{}
	result := &cobra.Command{
		Use:   "set-host <identifier>",
		Short: "Update a Dnsmasq host entry",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			console.Emit(c.SetHost(a[0], *hostRequest(f), &f.apply))
		},
	}
	registerHostFlags(result, f)

	return result
}
