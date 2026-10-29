package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createPrefix(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var site string
	var description string
	result := &cobra.Command{
		Use:   "create-prefix [cidr]",
		Short: "Create a NetBox IP prefix",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreatePrefix(
					context.Background(),
					arguments[0],
					site,
					description,
				),
			)
		},
	}
	result.Flags().StringVar(&site, "site", "", "site name (optional)")
	result.Flags().StringVar(
		&description,
		"description",
		"",
		"description (optional)",
	)

	return result
}
