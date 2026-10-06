package gojellyfin

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func libraries(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "libraries",
		Short: "List media libraries",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := x.Client.ListLibrariesWithResponse(context.Background())

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

			for _, l := range *r.JSON200 {
				kind := ""

				if l.CollectionType != nil {
					kind = *l.CollectionType
				}

				fmt.Printf("%s  %s  %s\n", l.Id, l.Name, kind)
			}
		},
	}
}
