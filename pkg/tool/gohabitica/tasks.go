package gohabitica

import (
	"fmt"
	"github.com/spf13/cobra"
)

func tasks(x *Context) *cobra.Command {
	var taskType string
	result := &cobra.Command{
		Use:   "tasks",
		Short: "List tasks",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(x.Client.Tasks(taskType))
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
