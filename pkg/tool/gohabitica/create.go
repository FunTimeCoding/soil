package gohabitica

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
	"github.com/spf13/cobra"
)

func create(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
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
			t.EmitValue(
				c.CreateTask(context.Background(), taskType, text, notes),
			)
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
