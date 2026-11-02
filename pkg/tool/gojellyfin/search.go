package gojellyfin

import (
	"context"
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
			p := &client.SearchItemsParams{Q: arguments[0]}

			if types != "" {
				p.Types = &types
			}

			if page > 0 {
				p.Page = &page
			}

			if perPage > 0 {
				p.PerPage = &perPage
			}

			x.Terminal.EmitValue(x.Client.SearchItems(context.Background(), p))
		},
	}
	result.Flags().StringVar(&types, "types", "", "comma-separated item types")
	result.Flags().IntVar(&page, "page", 0, "result page")
	result.Flags().IntVar(&perPage, "per-page", 0, "results per page")

	return result
}
