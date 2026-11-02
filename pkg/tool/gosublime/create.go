package gosublime

import (
	"context"
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

			x.Terminal.EmitValue(
				x.Client.CreateView(context.Background(), body),
			)
		},
	}
	result.Flags().StringVar(&syntax, "syntax", "", "syntax name")

	return result
}
