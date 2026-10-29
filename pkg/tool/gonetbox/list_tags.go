package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func listTags(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tags",
		Short: "List all NetBox tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.ListTags(context.Background()))
		},
	}
}
