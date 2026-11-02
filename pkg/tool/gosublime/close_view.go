package gosublime

import (
	"context"
	"github.com/spf13/cobra"
	"strconv"
)

func closeView(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "close <id>",
		Short: "Close a view",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])

			if e != nil {
				x.Terminal.Exitf("invalid id: %s\n", arguments[0])
			}

			x.Terminal.EmitDone(
				x.Client.CloseView(context.Background(), identifier),
			)
		},
	}
}
