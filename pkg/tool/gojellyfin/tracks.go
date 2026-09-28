package gojellyfin

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
	"os"
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
				fmt.Printf("%s  %s\n", v.Id, v.Name)
			}
		},
	}
}
