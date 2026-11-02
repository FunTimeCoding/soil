package gosublime

import (
	"context"
	"github.com/spf13/cobra"
)

func views(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "views",
		Short: "List open views",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			x.Terminal.EmitValue(x.Client.GetViews(context.Background()))
		},
	}
}
