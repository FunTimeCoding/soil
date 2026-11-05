package goband

import (
	"context"
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
			x.Terminal.EmitValue(x.Client.ListInstances(context.Background()))
		},
	}
}
