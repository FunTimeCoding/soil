package gopnsense

import (
	"fmt"
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
			fmt.Println(c.ReconfigureDnsmasq())
		},
	}
}
