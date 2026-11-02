package gosecret

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/spf13/cobra"
)

func showCommand(
	t *terminal.Terminal,
	c face.Client,
) *cobra.Command {
	return &cobra.Command{
		Use:   constant.ShowUsage,
		Short: "Show a secret's key names; values stay masked",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.ShowSecret(
					context.Background(),
					&client.ShowSecretParams{Path: arguments[0]},
				),
			)
		},
	}
}
