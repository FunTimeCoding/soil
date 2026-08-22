package gohabitica

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/spf13/cobra"
)

func create(x *Context) *cobra.Command {
	var taskType string
	var text string
	var notes string
	result := &cobra.Command{
		Use:   "create",
		Short: "Create a task",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(x.Client.CreateTask(taskType, text, notes))
			x.record(constant.CreateTask)
		},
	}
	result.Flags().StringVar(
		&taskType,
		"type",
		"",
		"Task type: habit, daily, todo, reward",
	)
	result.Flags().StringVar(&text, "text", "", "Task title")
	result.Flags().StringVar(&notes, "notes", "", "Task notes")
	result.MarkFlagsRequiredTogether("type", "text")

	return result
}
