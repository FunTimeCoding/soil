package gojellyfin

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/spf13/cobra"
)

func command(x *Context) *cobra.Command {
	var seek int64
	result := &cobra.Command{
		Use:   "command <session-identifier> <command>",
		Short: "Send a playback command (Pause, Unpause, Stop, Seek, ...)",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := client.PlaybackCommandJSONRequestBody{Command: arguments[1]}

			if seek > 0 {
				body.SeekPositionTicks = &seek
			}

			x.Terminal.EmitDone(
				x.Client.PlaybackCommand(
					context.Background(),
					arguments[0],
					body,
				),
			)
		},
	}
	result.Flags().Int64Var(&seek, "seek-ticks", 0, "seek position in ticks")

	return result
}
