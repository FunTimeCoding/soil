package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createLocation(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var site string
	result := &cobra.Command{
		Use:   "create-location [name]",
		Short: "Create a location within a site",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateLocation(context.Background(), arguments[0], site),
			)
		},
	}
	result.Flags().StringVar(&site, "site", "", "site name (required)")
	errors.PanicOnError(result.MarkFlagRequired("site"))

	return result
}
