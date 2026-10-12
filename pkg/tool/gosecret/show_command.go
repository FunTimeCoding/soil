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

func showCommand(t *terminal.Terminal) *cobra.Command {
	return &cobra.Command{
		Use:   constant.ShowUsage,
		Short: "Show a secret's key names; values stay masked",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := interfaceClient(t).ShowSecretWithResponse(
				context.Background(),
				&client.ShowSecretParams{Path: arguments[0]},
			)
			errors.PanicOnError(e)

			if r.JSON200 == nil {
				t.Reject(r.Status(), r.Body)

				return
			}

			for _, k := range r.JSON200.Keys {
				console.Line(k)
			}
		},
	}
}
