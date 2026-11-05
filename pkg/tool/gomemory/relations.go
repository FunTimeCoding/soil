package gomemory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemory/result"
	gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
)

func relations(x *Context) *cobra.Command {
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
				relationType = gomemorydConstant.UntypedFilter
			}

			if relationType != "" {
				params.Type = &relationType
			}

			if scope != "" {
				params.Scope = &scope
			}

			rows, e := x.Client.GetRelations(context.Background(), params)

			if e != nil {
				x.Terminal.EmitDone(e)

				return
			}

			x.Terminal.EmitValue(result.NewRelations(rows), nil)
		},
	}
	c.Flags().StringVar(&relationType, "type", "", "filter by relation type")
	c.Flags().BoolVar(&untyped, "untyped", false, "show only untyped relations")
	c.Flags().StringVar(
		&scope,
		"scope",
		"",
		"filter by scope on either endpoint (default for the default scope)",
	)

	return c
}
