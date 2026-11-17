package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/spf13/cobra"
)

func localContextFlag(c *cobra.Command) func() *map[string]any {
	var payload string
	c.Flags().StringVar(
		&payload,
		"local-context",
		"",
		"JSON object replacing the whole local context; {} clears it",
	)

	return func() *map[string]any {
		if !c.Flags().Changed("local-context") {
			return nil
		}

		result := map[string]any{}
		errors.PanicOnError(notation.Decode(payload, &result))

		return &result
	}
}
