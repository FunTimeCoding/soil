package gojellyfin

import (
	"context"
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
			x.Terminal.EmitValue(
				x.Client.GetItem(context.Background(), arguments[0]),
			)
		},
	}
}
