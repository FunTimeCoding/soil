package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func addHost(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	f := &hostFlags{}
	result := &cobra.Command{
		Use:   "add-host",
		Short: "Add a Dnsmasq host entry",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.AddHost(*hostRequest(f), &f.apply))
		},
	}
	registerHostFlags(result, f)

	return result
}
