package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/spf13/cobra"
)

func upgrade(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade <release> <source>",
		Short: "Show one upgrade with its hosts and the binary packages to upgrade on each",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitValue(c.GetUpgrade(context.Background(), a[0], a[1]))
		},
	}
}
