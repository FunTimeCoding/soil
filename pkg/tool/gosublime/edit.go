package gosublime

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func edit(x *Context) *cobra.Command {
	var old string
	var replacement string
	var all bool
	result := &cobra.Command{
		Use:   "edit <id>",
		Short: "Replace text in a view",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])

			if e != nil {
				x.Terminal.Exitf("invalid id: %s\n", arguments[0])
			}

			body := client.EditViewJSONRequestBody{
				OldString: old,
				NewString: replacement,
			}

			if all {
				body.ReplaceAll = &all
			}

			x.Terminal.EmitValue(
				x.Client.EditView(context.Background(), identifier, body),
			)
		},
	}
	result.Flags().StringVar(&old, "old", "", "text to replace")
	result.Flags().StringVar(&replacement, "new", "", "replacement text")
	result.Flags().BoolVar(&all, "all", false, "replace all occurrences")
	errors.PanicOnError(result.MarkFlagRequired("old"))
	errors.PanicOnError(result.MarkFlagRequired("new"))

	return result
}
