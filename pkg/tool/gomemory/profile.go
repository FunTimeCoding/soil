package gomemory

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/spf13/cobra"
	"io"
)

func profile(l **client.Client) *cobra.Command {
	var topic string
	var detail bool
	c := &cobra.Command{
		Use:   "profile",
		Short: "Load memory profile",
		Run: func(_ *cobra.Command, _ []string) {
			params := &client.GetProfileParams{}

			if topic != "" {
				params.Topic = &topic
			}

			if detail {
				params.Detail = &detail
			}

			r, e := (*l).GetProfile(
				context.Background(),
				params,
			)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			body, e := io.ReadAll(r.Body)
			errors.PanicOnError(e)
			fmt.Print(string(body))
		},
	}
	c.Flags().StringVar(
		&topic,
		"topic",
		"",
		"session topic for relevance matching",
	)
	c.Flags().BoolVar(
		&detail,
		"detail",
		false,
		"include token budget details",
	)

	return c
}
