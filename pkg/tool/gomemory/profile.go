package gomemory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
	"strings"
)

func profile(x *Context) *cobra.Command {
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

			r, e := x.Client.GetProfile(context.Background(), params)

			if e != nil || detail {
				x.Terminal.EmitValue(r, e)

				return
			}

			x.Terminal.Line(strings.TrimSuffix(r.Text, "\n"))
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
