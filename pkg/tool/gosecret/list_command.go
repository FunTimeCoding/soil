package gosecret

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/spf13/cobra"
	"net/http"
)

func listCommand(
	t *terminal.Terminal,
	h *http.Client,
) *cobra.Command {
	return &cobra.Command{
		Use:   constant.ListUsage,
		Short: "List secrets and folders below a path; folders end in a slash",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			p := &client.ListSecretsParams{}

			if len(arguments) > 0 {
				p.Path = &arguments[0]
			}

			r, e := interfaceClient(t, h).ListSecretsWithResponse(
				context.Background(),
				p,
			)
			errors.PanicOnError(e)

			if r.JSON200 == nil {
				t.Reject(r.Status(), r.Body)

				return
			}

			for _, entry := range r.JSON200.Entries {
				console.Line(entry)
			}
		},
	}
}
