package gojellyfin

import (
	"context"
	"github.com/spf13/cobra"
)

func episodes(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "episodes <series-identifier>",
		Short: "List episodes of a series",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			x.Terminal.EmitValue(
				x.Client.GetEpisodes(context.Background(), arguments[0]),
			)
		},
	}
}
