package gojellyfin

import (
	"context"
	"fmt"
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
			r, e := x.Client.GetEpisodesWithResponse(
				context.Background(),
				arguments[0],
			)

			if e != nil {
				x.Terminal.Exitf("error: %v\n", e)
			}

			if r.JSON200 == nil {
				x.Terminal.Exitf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
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
