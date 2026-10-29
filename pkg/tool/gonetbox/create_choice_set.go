package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createChoiceSet(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var choices []string
	result := &cobra.Command{
		Use:   "create-choice-set [name]",
		Short: "Create a custom field choice set",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateChoiceSet(context.Background(), arguments[0], choices),
			)
		},
	}
	result.Flags().StringSliceVar(
		&choices,
		"choice",
		nil,
		"choice values (repeatable or comma separated)",
	)

	return result
}
