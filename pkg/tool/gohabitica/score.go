package gohabitica

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
	"github.com/spf13/cobra"
)

func score(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var direction string
	result := &cobra.Command{
		Use:   "score [identifier]",
		Short: "Score a task",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.ScoreTask(context.Background(), arguments[0], direction),
			)
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
