package gomemory

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
)

func statistic(l **client.Client) *cobra.Command {
	var scope string
	c := &cobra.Command{
		Use:   "statistic",
		Short: "Token statistic per memory, heaviest first",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			params := &client.GetTokensParams{}

			if scope != "" {
				params.Scope = &scope
			}

			r, e := (*l).GetTokens(context.Background(), params)
			errors.PanicOnError(e)
			parsed, f := client.ParseGetTokensResponse(r)
			errors.PanicOnError(f)

			if parsed.JSON200 == nil {
				fmt.Print(string(parsed.Body))

				return
			}

			fmt.Print(statisticTable(parsed.JSON200))
		},
	}
	c.Flags().StringVar(&scope, "scope", "", "memory scope")

	return c
}
