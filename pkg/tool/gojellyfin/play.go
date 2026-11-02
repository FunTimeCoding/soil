package gojellyfin

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/spf13/cobra"
)

func play(x *Context) *cobra.Command {
	var playCommand string
	var start int64
	result := &cobra.Command{
		Use:   "play <session-identifier> <item-identifier>...",
		Short: "Play items on a session",
		Args:  cobra.MinimumNArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := client.PlayJSONRequestBody{ItemIds: arguments[1:]}

			if playCommand != "" {
				body.PlayCommand = &playCommand
			}

			if start > 0 {
				body.StartPositionTicks = &start
			}

			x.Terminal.EmitDone(
				x.Client.Play(context.Background(), arguments[0], body),
			)
		},
	}
	result.Flags().StringVar(
		&playCommand,
		"command",
		"",
		"play command (default PlayNow)",
	)
	result.Flags().Int64Var(&start, "start-ticks", 0, "start position in ticks")

	return result
}
