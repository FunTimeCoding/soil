package gohabitica

import (
	"fmt"
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
			fmt.Println(c.Tasks(taskType))
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
