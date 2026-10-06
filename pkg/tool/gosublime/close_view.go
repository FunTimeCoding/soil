package gosublime

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/spf13/cobra"
	"strconv"
)

func closeView(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "close <id>",
		Short: "Close a view",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])

			if e != nil {
				x.Terminal.Exitf("invalid id: %s\n", arguments[0])
			}

			r, f := x.Client.CloseViewWithResponse(
				context.Background(),
				identifier,
			)

			if f != nil {
				x.Terminal.Exitf("error: %v\n", f)
			}

			if r.HTTPResponse.StatusCode != 204 {
				x.Terminal.Exitf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
			}

			console.Format("closed view %d\n", identifier)
		},
	}
}
