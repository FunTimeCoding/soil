package gosecret

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/spf13/cobra"
	"net/http"
)

func statusCommand(
	t *terminal.Terminal,
	h *http.Client,
) *cobra.Command {
	return &cobra.Command{
		Use:   constant.StatusUsage,
		Short: "Show whether the secret store is initialized and sealed",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := interfaceClient(t, h).GetStatusWithResponse(
				context.Background(),
			)
			errors.PanicOnError(e)

			if r.JSON200 == nil {
				t.Reject(r.Status(), r.Body)

				return
			}

			console.Format(
				"initialized: %v\nsealed: %v\nversion: %s\n",
				r.JSON200.Initialized,
				r.JSON200.Sealed,
				r.JSON200.Version,
			)
		},
	}
}
