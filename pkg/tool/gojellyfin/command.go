package gojellyfin

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/spf13/cobra"
	"os"
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

			r, e := x.Client.PlaybackCommandWithResponse(
				context.Background(),
				arguments[0],
				body,
			)

			if e != nil {
				errors.Printf("error: %v\n", e)
				os.Exit(1)
			}

			if r.HTTPResponse.StatusCode != 204 {
				errors.Printf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
				os.Exit(1)
			}

			fmt.Printf("sent %s\n", arguments[1])
		},
	}
	result.Flags().Int64Var(&seek, "seek-ticks", 0, "seek position in ticks")

	return result
}
