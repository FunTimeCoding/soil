package gojellyfin

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
	"os"
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

			fmt.Println(string(r.Body))
		},
	}
}
