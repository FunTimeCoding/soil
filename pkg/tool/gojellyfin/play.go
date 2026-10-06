package gojellyfin

import (
	"context"
	"fmt"
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

			r, e := x.Client.PlayWithResponse(
				context.Background(),
				arguments[0],
				body,
			)

			if e != nil {
				x.Terminal.Exitf("error: %v\n", e)
			}

			if r.HTTPResponse.StatusCode != 204 {
				x.Terminal.Exitf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
			}

			fmt.Println("playback started")
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
