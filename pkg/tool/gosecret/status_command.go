package gosecret

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/spf13/cobra"
)

func statusCommand(
	t *terminal.Terminal,
	c face.Client,
) *cobra.Command {
	return &cobra.Command{
		Use:   constant.StatusUsage,
		Short: "Show whether the secret store is initialized and sealed",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.GetStatus(context.Background()))
		},
	}
}
