package gosublime

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/spf13/cobra"
)

func create(x *Context) *cobra.Command {
	var syntax string
	result := &cobra.Command{
		Use:   "create <title> <content>",
		Short: "Create a scratch view",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := client.CreateViewJSONRequestBody{
				Title:   arguments[0],
				Content: arguments[1],
			}

			if syntax != "" {
				body.Syntax = &syntax
			}

			r, e := x.Client.CreateViewWithResponse(context.Background(), body)

			if e != nil {
				x.Terminal.Exitf("error: %v\n", e)
			}

			if r.JSON200 == nil {
				x.Terminal.Exitf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
			}

			console.Format("created view %d\n", r.JSON200.ViewId)
		},
	}
	result.Flags().StringVar(&syntax, "syntax", "", "syntax name")

	return result
}
