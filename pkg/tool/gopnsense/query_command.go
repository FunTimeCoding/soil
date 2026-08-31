package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/spf13/cobra"
)

func queryCommand(
	use string,
	short string,
	call func(query *string) *response.Response,
) *cobra.Command {
	var query string
	result := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var q *string

			if query != "" {
				q = &query
			}

			console.Emit(call(q))
		},
	}
	result.Flags().StringVar(&query, "query", "", "search phrase")

	return result
}
