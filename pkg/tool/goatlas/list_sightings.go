package goatlas

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goatlas/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
	"github.com/spf13/cobra"
)

func listSightings(x *Context) *cobra.Command {
	var unclaimed bool
	var place string
	result := &cobra.Command{
		Use:   "list-sightings",
		Short: "List hosts seen on the network and what they resolve to",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			p := &client.ListSightingsParams{}

			if unclaimed {
				p.Unclaimed = &unclaimed
			}

			if place != "" {
				p.Place = &place
			}

			r, e := x.Client.ListSightingsWithResponse(context.Background(), p)

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
					"%-16s %-24s %-12s %s\n",
					v.Address,
					sightingName(v),
					sightingPlace(v),
					time.FormatCompact(v.Seen),
				)
			}
		},
	}
	result.Flags().BoolVar(
		&unclaimed,
		constant.UnclaimedArgument,
		false,
		constant.UnclaimedUsage,
	)
	result.Flags().StringVar(
		&place,
		constant.PlaceArgument,
		"",
		constant.PlaceUsage,
	)

	return result
}
