package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/spf13/cobra"
)

func embed(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "embed",
		Short: "Generate embeddings for indexed documents",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.PostEmbed(context.Background()))
		},
	}
}
