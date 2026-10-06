package gojellyfin

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/spf13/cobra"
)

func search(x *Context) *cobra.Command {
	var types string
	var page int
	var perPage int
	result := &cobra.Command{
		Use:   "search <term>",
		Short: "Search items",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			params := &client.SearchItemsParams{Q: arguments[0]}

			if types != "" {
				params.Types = &types
			}

			if page > 0 {
				params.Page = &page
			}

			if perPage > 0 {
				params.PerPage = &perPage
			}

			r, e := x.Client.SearchItemsWithResponse(
				context.Background(),
				params,
			)

			if e != nil {
				x.Terminal.Exitf("error: %v\n", e)
			}

			if r.JSON200 == nil {
				x.Terminal.Exitf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
			}

			for _, v := range r.JSON200.Items {
				fmt.Printf("%s  [%s]  %s\n", v.Id, v.Type, v.Name)
			}

			fmt.Printf("page %d, %d total\n", r.JSON200.Page, r.JSON200.Total)
		},
	}
	result.Flags().StringVar(&types, "types", "", "comma-separated item types")
	result.Flags().IntVar(&page, "page", 0, "result page")
	result.Flags().IntVar(&perPage, "per-page", 0, "results per page")

	return result
}
