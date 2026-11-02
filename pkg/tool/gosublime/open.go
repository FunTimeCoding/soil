package gosublime

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/spf13/cobra"
)

func open(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "open <path>",
		Short: "Open a file in Sublime",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			x.Terminal.EmitValue(
				x.Client.OpenFile(
					context.Background(),
					client.OpenFileJSONRequestBody{FilePath: arguments[0]},
				),
			)
		},
	}
}
