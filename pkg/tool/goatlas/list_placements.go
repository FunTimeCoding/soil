package goatlas

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlas/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/spf13/cobra"
)

func listPlacements(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var place string
	var source string
	var packageName string
	result := &cobra.Command{
		Use:   "list-placements",
		Short: "List placements, newest attestation per line",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(
				c.ListPlacements(
					context.Background(),
					place,
					source,
					packageName,
				),
			)
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
	result.Flags().StringVar(
		&packageName,
		constant.PackageArgument,
		"",
		constant.PackageUsage,
	)

	return result
}
