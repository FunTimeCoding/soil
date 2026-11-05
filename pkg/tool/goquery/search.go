package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func search(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
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

			t.EmitValue(c.GetSearch(context.Background(), params))
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
