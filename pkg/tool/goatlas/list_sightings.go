package goatlas

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlas/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/spf13/cobra"
)

func listSightings(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
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
			t.EmitValue(c.ListSightings(context.Background(), place, unclaimed))
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
