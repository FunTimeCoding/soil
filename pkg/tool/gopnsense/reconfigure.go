package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func reconfigure(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "reconfigure",
		Short: "Apply pending Dnsmasq configuration",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ReconfigureDnsmasq())
		},
	}
}
