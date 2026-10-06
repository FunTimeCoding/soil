package gojellyfin

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func tracks(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "tracks <album-identifier>",
		Short: "List tracks of an album",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := x.Client.GetTracksWithResponse(
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
				fmt.Printf("%s  %s\n", v.Id, v.Name)
			}
		},
	}
}
