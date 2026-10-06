package gosublime

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func read(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "read <id>",
		Short: "Read a view's full text",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])

			if e != nil {
				x.Terminal.Exitf("invalid id: %s\n", arguments[0])
			}

			r, e := x.Client.GetViewWithResponse(
				context.Background(),
				identifier,
			)

			if e != nil {
				x.Terminal.Exitf("error: %v\n", e)
			}

			if r.JSON200 == nil {
				x.Terminal.Exitf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
			}

			if r.JSON200.Text != nil {
				fmt.Print(*r.JSON200.Text)
			}
		},
	}
}
