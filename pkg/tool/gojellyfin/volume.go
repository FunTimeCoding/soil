package gojellyfin

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func volume(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "volume <session-identifier> <level>",
		Short: "Set a session's volume (0-100)",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			level, e := strconv.Atoi(arguments[1])

			if e != nil {
				x.Terminal.Exitf("invalid level: %s\n", arguments[1])
			}

			r, f := x.Client.SetVolumeWithResponse(
				context.Background(),
				arguments[0],
				client.SetVolumeJSONRequestBody{Level: level},
			)

			if f != nil {
				x.Terminal.Exitf("error: %v\n", f)
			}

			if r.HTTPResponse.StatusCode != 204 {
				x.Terminal.Exitf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
			}

			fmt.Printf("volume set to %d\n", level)
		},
	}
}
