package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/spf13/cobra"
)

func listContexts(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "context-list",
		Short: "List all contexts",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			contexts, e := c.GetContext(context.Background())
			t.EmitDone(e)

			for _, v := range contexts {
				console.Format(
					"%s  %s  %s\n",
					v.Collection,
					v.PathPrefix,
					v.Description,
				)
			}
		},
	}
}
