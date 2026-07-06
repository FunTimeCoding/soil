package goquery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
	"strings"
)

func search(c *client.Client) *cobra.Command {
	var limit int
	var collection string
	var keywordOnly bool
	var detail bool
	var sourceType string
	var meta []string
	result := &cobra.Command{
		Use:   "search [query]",
		Short: "Search indexed documents",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			mode := client.Hybrid

			if keywordOnly {
				mode = client.Keyword
			}

			full := detail
			params := &client.GetSearchParams{
				Query: arguments[0],
				Limit: &limit,
				Mode:  &mode,
				Full:  &full,
			}

			if collection != "" {
				params.Collection = &collection
			}

			if sourceType != "" {
				params.SourceType = &sourceType
			}

			metadata := parseMetadata(meta)

			if metadata != nil {
				params.Metadata = &metadata
			}

			r, e := c.GetSearch(context.Background(), params)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var outcome client.SearchOutcome
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&outcome))

			for _, v := range outcome.Results {
				fmt.Printf(
					"%.4f  %s  %s",
					v.Score,
					v.Path,
					v.Title,
				)

				if v.Metadata != nil && len(*v.Metadata) > 0 {
					fmt.Printf("  %s", formatMetadata(*v.Metadata))
				}

				fmt.Println()

				if detail && v.Body != nil && *v.Body != "" {
					for _, line := range strings.Split(*v.Body, "\n") {
						fmt.Printf("    %s\n", line)
					}
				}
			}

			printFacets(outcome.Facets)
		},
	}
	result.Flags().IntVar(&limit, "limit", 10, "Maximum number of results")
	result.Flags().StringVar(
		&collection,
		"collection",
		"",
		"Restrict search to a collection",
	)
	result.Flags().BoolVar(
		&keywordOnly,
		"keyword",
		false,
		"Use keyword search only",
	)
	result.Flags().BoolVar(
		&detail,
		"detail",
		false,
		"Show metadata and document body",
	)
	result.Flags().StringVar(
		&sourceType,
		"source-type",
		"",
		"Filter by source type",
	)
	result.Flags().StringArrayVar(
		&meta,
		"meta",
		nil,
		"Filter by metadata (key=value, repeatable)",
	)

	return result
}
