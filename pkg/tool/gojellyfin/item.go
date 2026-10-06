package gojellyfin

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func item(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "item <identifier>",
		Short: "Get an item's detail",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := x.Client.GetItemWithResponse(
				context.Background(),
				arguments[0],
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

			fmt.Println(string(r.Body))
		},
	}
}
