package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func tasks(c *client.Client) *cobra.Command {
	var taskType string
	result := &cobra.Command{
		Use:   "tasks",
		Short: "List tasks",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.Tasks(taskType))
		},
	}
	result.Flags().StringVar(
		&taskType,
		"type",
		"",
		"Task type: habits, dailys, todos, rewards",
	)

	return result
}
