package gomemory

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
	"io"
	"os"
	"text/tabwriter"
)

func relations(l **client.Client) *cobra.Command {
	var relationType string
	var untyped bool
	var scope string
	c := &cobra.Command{
		Use:   "relations",
		Short: "List the relation graph with names, scopes, and types",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			params := &client.GetRelationsParams{}

			if untyped {
				relationType = constant.UntypedFilter
			}

			if relationType != "" {
				params.Type = &relationType
			}

			if scope != "" {
				params.Scope = &scope
			}

			r, e := (*l).GetRelations(
				context.Background(),
				params,
			)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			body, e := io.ReadAll(r.Body)
			errors.PanicOnError(e)
			var rows []client.Relation
			errors.PanicOnError(json.Unmarshal(body, &rows))
			w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
			_, e = fmt.Fprintln(
				w,
				"SOURCE\tSCOPE\tTYPE\tTARGET\tSCOPE",
			)
			errors.PanicOnError(e)

			for _, row := range rows {
				_, e = fmt.Fprintf(
					w,
					"%s\t%s\t%s\t%s\t%s\n",
					row.SourceName,
					scopeLabel(row.SourceScope),
					typeLabel(row.Type),
					row.TargetName,
					scopeLabel(row.TargetScope),
				)
				errors.PanicOnError(e)
			}

			errors.PanicFlush(w)
			fmt.Printf("%d relations\n", len(rows))
		},
	}
	c.Flags().StringVar(
		&relationType,
		"type",
		"",
		"filter by relation type",
	)
	c.Flags().BoolVar(
		&untyped,
		"untyped",
		false,
		"show only untyped relations",
	)
	c.Flags().StringVar(
		&scope,
		"scope",
		"",
		"filter by scope on either endpoint (default for the default scope)",
	)

	return c
}
