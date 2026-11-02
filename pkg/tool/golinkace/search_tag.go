package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
)

func searchTag(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "search-tag <query>",
		Short: "Search tags by name",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			entityType := "tag"
			c.Terminal().EmitValue(
				c.Client().Search(
					context.Background(),
					&client.SearchParams{Query: arguments[0], Type: &entityType},
				),
			)
		},
	}
}
