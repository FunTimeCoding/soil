package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/spf13/cobra"
)

func scores(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "scores",
		Short: "Show how far the score mirror has come and when it last caught up",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.GetScoreStatus(context.Background()))
		},
	}
}
