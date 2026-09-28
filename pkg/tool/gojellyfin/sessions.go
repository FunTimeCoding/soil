package gojellyfin

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
	"os"
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
			r, e := x.Client.ListSessionsWithResponse(context.Background())

			if e != nil {
				errors.Printf("error: %v\n", e)
				os.Exit(1)
			}

			if r.JSON200 == nil {
				errors.Printf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
				os.Exit(1)
			}

			for _, v := range *r.JSON200 {
				remote := " "

				if v.SupportsRemote {
					remote = "*"
				}

				playing := ""

				if v.NowPlayingName != nil {
					playing = *v.NowPlayingName
				}

				fmt.Printf(
					"%s %s %s (%s)  %s  %s\n",
					remote,
					v.Id,
					v.DeviceName,
					v.Client,
					v.PlayState,
					playing,
				)
			}
		},
	}
}
