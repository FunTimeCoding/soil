package gojellyfin

import (
	"context"
	"github.com/spf13/cobra"
)

func libraries(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "libraries",
		Short: "List media libraries",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			x.Terminal.EmitValue(x.Client.ListLibraries(context.Background()))
		},
	}
}
