package gosublime

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func save(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "save <id> [path]",
		Short: "Save a view, optionally to a new path",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])

			if e != nil {
				x.Terminal.Exitf("invalid id: %s\n", arguments[0])
			}

			var body client.SaveViewJSONRequestBody

			if len(arguments) > 1 {
				body.FilePath = &arguments[1]
			}

			x.Terminal.EmitDone(
				x.Client.SaveView(context.Background(), identifier, body),
			)
		},
	}
}
