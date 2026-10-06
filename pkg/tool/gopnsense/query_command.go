package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/spf13/cobra"
)

func queryCommand(
	use string,
	short string,
	call func(query *string) *response.Response,
	t *terminal.Terminal,
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

			t.Emit(call(q))
		},
	}
	result.Flags().StringVar(&query, "query", "", "search phrase")

	return result
}
