package gosecret

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/spf13/cobra"
)

func revealCommand(
	t *terminal.Terminal,
	c face.Client,
) *cobra.Command {
	return &cobra.Command{
		Use:   constant.RevealUsage,
		Short: "Print one key's value and nothing else",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			v, e := c.RevealSecret(
				context.Background(),
				&client.RevealSecretParams{
					Path: arguments[0],
					Key:  arguments[1],
				},
			)

			if e != nil {
				t.Exitln(e)

				return
			}

			console.Line(v.Value)
		},
	}
}
