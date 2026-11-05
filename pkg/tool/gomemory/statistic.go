package gomemory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
)

func statistic(x *Context) *cobra.Command {
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

			x.Terminal.EmitValue(
				x.Client.GetTokens(context.Background(), params),
			)
		},
	}
	c.Flags().StringVar(&scope, "scope", "", "memory scope")

	return c
}
