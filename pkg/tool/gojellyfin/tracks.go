package gojellyfin

import (
	"context"
	"github.com/spf13/cobra"
)

func tracks(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "tracks <album-identifier>",
		Short: "List tracks of an album",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			x.Terminal.EmitValue(
				x.Client.GetTracks(context.Background(), arguments[0]),
			)
		},
	}
}
