package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createCustomField(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var choiceSet string
	var label string
	result := &cobra.Command{
		Use:   "create-custom-field [name]",
		Short: "Create a multiple-selection custom field on devices and virtual machines",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateCustomField(
					context.Background(),
					arguments[0],
					label,
					choiceSet,
				),
			)
		},
	}
	result.Flags().StringVar(
		&choiceSet,
		"choice-set",
		"",
		"choice set the field selects from (required)",
	)
	result.Flags().StringVar(&label, constant.Label, "", "display label")

	return result
}
