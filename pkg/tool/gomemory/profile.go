package gomemory

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
)

func profile(l **client.Client) *cobra.Command {
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
			errors.PanicOnError(e)
			parsed, f := client.ParseGetProfileResponse(r)
			errors.PanicOnError(f)

			if parsed.JSON200 == nil {
				fmt.Print(string(parsed.Body))

				return
			}

			fmt.Print(parsed.JSON200.Text)

			if parsed.JSON200.Detail != nil {
				fmt.Print(budgetTable(parsed.JSON200.Detail, parsed.JSON200))
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
