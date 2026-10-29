package gopnsense

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/spf13/cobra"
)

func addHost(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	f := &HostFlags{}
	result := &cobra.Command{
		Use:   "add-host",
		Short: "Add a Dnsmasq host entry",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(
				c.AddHost(context.Background(), *hostRequest(f), f.apply),
			)
		},
	}
	registerHostFlags(result, f)

	return result
}
