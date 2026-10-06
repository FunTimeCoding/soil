package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createPrefix(
	c *client.Client,
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
			t.Emit(c.CreatePrefix(arguments[0], site, description))
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
