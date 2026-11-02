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
			instances, e := x.Client.ListInstances(context.Background())

			if e != nil {
				x.Terminal.Exitln(e)

				return
			}

			for _, v := range instances {
				x.Terminal.Line(fmt.Sprintf("%-12s %s", v.Name, v.Host))
			}
		},
	}
}
