package gopnsense

import (
	"fmt"
	"github.com/spf13/cobra"
)

func queryCommand(
	use string,
	short string,
	call func(query *string) string,
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

			fmt.Println(call(q))
		},
	}
	result.Flags().StringVar(&query, "query", "", "search phrase")

	return result
}
