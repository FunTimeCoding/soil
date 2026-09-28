package gojellyfin

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
	"os"
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
			r, e := x.Client.GetEpisodesWithResponse(
				context.Background(),
				arguments[0],
			)

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
				season := 0
				episode := 0

				if v.SeasonNumber != nil {
					season = *v.SeasonNumber
				}

				if v.EpisodeNumber != nil {
					episode = *v.EpisodeNumber
				}

				fmt.Printf(
					"%s  S%02dE%02d  %s\n",
					v.Id,
					season,
					episode,
					v.Name,
				)
			}
		},
	}
}
