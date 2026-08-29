package gopnsense

import (
	"fmt"
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
			fmt.Println(c.AddHost(*hostRequest(f), &f.apply))
		},
	}
	registerHostFlags(result, f)

	return result
}
