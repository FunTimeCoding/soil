package gojellyfin

import (
	"context"
	"github.com/spf13/cobra"
)

func sessions(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "sessions",
		Short: "List active playback sessions",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			x.Terminal.EmitValue(x.Client.ListSessions(context.Background()))
		},
	}
}
