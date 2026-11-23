package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/spf13/cobra"
)

func system(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "system <name>",
		Short: "Show one system's coverage, kernel and findings",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitValue(c.GetSystemFindings(context.Background(), a[0]))
		},
	}
}
