package gopnsense

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/spf13/cobra"
)

func reconfigure(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "reconfigure",
		Short: "Apply pending Dnsmasq configuration",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitDone(c.ReconfigureDnsmasq(context.Background()))
		},
	}
}
