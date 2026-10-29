package gopnsense

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/spf13/cobra"
)

func setHost(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	f := &HostFlags{}
	result := &cobra.Command{
		Use:   "set-host <identifier>",
		Short: "Update a Dnsmasq host entry",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitDone(
				c.SetHost(context.Background(), a[0], *hostRequest(f), f.apply),
			)
		},
	}
	registerHostFlags(result, f)

	return result
}
