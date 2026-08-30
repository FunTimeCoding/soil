package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func reconfigure(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "reconfigure",
		Short: "Apply pending Dnsmasq configuration",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ReconfigureDnsmasq())
		},
	}
}
