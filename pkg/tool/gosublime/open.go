package gosublime

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/spf13/cobra"
	"os"
)

func open(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "open <path>",
		Short: "Open a file in Sublime",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := x.Client.OpenFileWithResponse(
				context.Background(),
				client.OpenFileJSONRequestBody{FilePath: arguments[0]},
			)

			if e != nil {
				errors.Printf("error: %v\n", e)
				os.Exit(1)
			}

			if r.JSON200 == nil {
				errors.Printf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
				os.Exit(1)
			}

			console.Format("opened view %d\n", r.JSON200.ViewId)
		},
	}
}
