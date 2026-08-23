package gosublime

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/spf13/cobra"
	"os"
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
				errors.Printf("invalid id: %s\n", arguments[0])
				os.Exit(1)
			}

			body := client.EditViewJSONRequestBody{
				OldString: old,
				NewString: replacement,
			}

			if all {
				body.ReplaceAll = &all
			}

			r, f := x.Client.EditViewWithResponse(
				context.Background(),
				identifier,
				body,
			)

			if f != nil {
				errors.Printf("error: %v\n", f)
				os.Exit(1)
			}

			if r.JSON200 == nil {
				errors.Printf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
				os.Exit(1)
			}

			fmt.Printf("edited view %d\n", identifier)
		},
	}
	result.Flags().StringVar(&old, "old", "", "text to replace")
	result.Flags().StringVar(&replacement, "new", "", "replacement text")
	result.Flags().BoolVar(&all, "all", false, "replace all occurrences")
	errors.PanicOnError(result.MarkFlagRequired("old"))
	errors.PanicOnError(result.MarkFlagRequired("new"))

	return result
}
