package goatlas

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goatlas/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
	"github.com/spf13/cobra"
)

func listPlacements(x *Context) *cobra.Command {
	var place string
	var source string
	result := &cobra.Command{
		Use:   "list-placements",
		Short: "List placements, newest attestation per line",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			p := &client.ListPlacementsParams{}

			if place != "" {
				p.Place = &place
			}

			if source != "" {
				p.Source = &source
			}

			r, e := x.Client.ListPlacementsWithResponse(context.Background(), p)

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
				fmt.Printf(
					"%-12s %-28s %-12s %s\n",
					placeName(v),
					qualifiedName(v),
					v.Source,
					time.FormatCompact(v.Seen),
				)
			}
		},
	}
	result.Flags().StringVar(
		&place,
		constant.PlaceArgument,
		"",
		constant.PlaceUsage,
	)
	result.Flags().StringVar(
		&source,
		constant.SourceArgument,
		"",
		constant.SourceUsage,
	)

	return result
}
