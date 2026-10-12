package goband

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func listInstances(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list-instances",
		Short: "List the configured management engines",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := x.Client.ListInstancesWithResponse(context.Background())

			if e != nil {
				x.Terminal.Exitf("error: %v\n", e)
			}

			if r.JSON200 == nil {
				x.Terminal.Reject(r.Status(), r.Body)

				return
			}

			for _, v := range *r.JSON200 {
				x.Terminal.Line(fmt.Sprintf("%-12s %s", v.Name, v.Host))
			}
		},
	}
}
