package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/spf13/cobra"
	"strings"
)

func customFieldFlag(c *cobra.Command) func() *map[string]any {
	var assignments []string
	c.Flags().StringArrayVar(
		&assignments,
		"custom-field",
		nil,
		"field=value,value (repeatable); field= clears it",
	)

	return func() *map[string]any {
		if len(assignments) == 0 {
			return nil
		}

		result := map[string]any{}

		for _, a := range assignments {
			field, values, okay := strings.Cut(a, "=")

			if !okay || field == "" {
				errors.PanicOnError(
					unexpected.Format("not field=values: %s", a),
				)
			}

			selected := []string{}

			if values != "" {
				selected = split.Comma(values)
			}

			result[field] = selected
		}

		return &result
	}
}
