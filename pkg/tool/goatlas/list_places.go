package goatlas

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func listPlaces(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list-places",
		Short: "List devices and virtual machines by how much they carry",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := x.Client.ListPlacesWithResponse(context.Background())

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
				fmt.Printf("%-12s %4d  %s\n", v.Name, v.Count, v.Kind)
			}
		},
	}
}
