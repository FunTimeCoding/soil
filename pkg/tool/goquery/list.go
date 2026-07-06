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

func list(c *client.Client) *cobra.Command {
	var limit int
	var offset int
	var detail bool
	var sourceType string
	var meta []string
	result := &cobra.Command{
		Use:   "list [collection]",
		Short: "List documents in a collection",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			full := detail
			params := &client.GetListParams{
				Collection: arguments[0],
				Full:       &full,
			}

			if limit > 0 {
				params.Limit = &limit
			}

			if offset > 0 {
				params.Offset = &offset
			}

			if sourceType != "" {
				params.SourceType = &sourceType
			}

			metadata := parseMetadata(meta)

			if metadata != nil {
				params.Metadata = &metadata
			}

			r, e := c.GetList(context.Background(), params)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var outcome client.ListOutcome
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&outcome))

			for _, v := range outcome.Results {
				fmt.Printf("%s  %s", v.Path, v.Title)

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
	result.Flags().IntVar(&offset, "offset", 0, "Number of results to skip")
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
