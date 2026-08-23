package gosublime

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
	"os"
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
				errors.Printf("invalid id: %s\n", arguments[0])
				os.Exit(1)
			}

			r, f := x.Client.CloseViewWithResponse(
				context.Background(),
				identifier,
			)

			if f != nil {
				errors.Printf("error: %v\n", f)
				os.Exit(1)
			}

			if r.HTTPResponse.StatusCode != 204 {
				errors.Printf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
				os.Exit(1)
			}

			fmt.Printf("closed view %d\n", identifier)
		},
	}
}
