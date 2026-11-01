package gomemory

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
)

func statistic(
	l *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
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

			if e != nil {
				t.Exitln(e)

				return
			}

			fmt.Print(statisticTable(r))
		},
	}
	c.Flags().StringVar(&scope, "scope", "", "memory scope")

	return c
}
