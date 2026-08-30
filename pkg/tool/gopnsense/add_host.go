package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func addHost(c *client.Client) *cobra.Command {
	f := &hostFlags{}
	result := &cobra.Command{
		Use:   "add-host",
		Short: "Add a Dnsmasq host entry",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.AddHost(*hostRequest(f), &f.apply))
		},
	}
	registerHostFlags(result, f)

	return result
}
