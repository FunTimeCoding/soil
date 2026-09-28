package goatlas

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
	"os"
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
				fmt.Printf("%-12s %4d  %s\n", v.Name, v.Count, v.Kind)
			}
		},
	}
}
