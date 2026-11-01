package gomemory

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
)

func profile(
	l *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var topic string
	var detail bool
	c := &cobra.Command{
		Use:   "profile",
		Short: "Load memory profile",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			params := &client.GetProfileParams{}

			if topic != "" {
				params.Topic = &topic
			}

			if detail {
				params.Detail = &detail
			}

			r, e := (*l).GetProfile(context.Background(), params)

			if e != nil {
				t.Exitln(e)

				return
			}

			fmt.Print(r.Text)

			if r.Detail != nil {
				fmt.Print(budgetTable(r.Detail, r))
			}
		},
	}
	c.Flags().StringVar(
		&topic,
		"topic",
		"",
		"session topic for relevance matching",
	)
	c.Flags().BoolVar(&detail, "detail", false, "include token budget details")

	return c
}
