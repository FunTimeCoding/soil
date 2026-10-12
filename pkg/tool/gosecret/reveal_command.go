package gosecret

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/spf13/cobra"
)

func revealCommand(t *terminal.Terminal) *cobra.Command {
	return &cobra.Command{
		Use:   constant.RevealUsage,
		Short: "Print one key's value and nothing else",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := interfaceClient(t).RevealSecretWithResponse(
				context.Background(),
				&client.RevealSecretParams{
					Path: arguments[0],
					Key:  arguments[1],
				},
			)
			errors.PanicOnError(e)

			if r.JSON200 == nil {
				t.Reject(r.Status(), r.Body)

				return
			}

			console.Line(r.JSON200.Value)
		},
	}
}
