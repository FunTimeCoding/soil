package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func score(c *client.Client) *cobra.Command {
	var direction string
	result := &cobra.Command{
		Use:   "score [identifier]",
		Short: "Score a task",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.Score(arguments[0], direction))
		},
	}
	result.Flags().StringVar(
		&direction,
		"direction",
		"up",
		"Score direction: up or down",
	)

	return result
}
