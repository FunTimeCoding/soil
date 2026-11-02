package gosublime

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func read(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "read <id>",
		Short: "Read a view's full text",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])

			if e != nil {
				x.Terminal.Exitf("invalid id: %s\n", arguments[0])
			}

			v, f := x.Client.GetView(context.Background(), identifier)

			if f != nil {
				x.Terminal.Exitln(f)
			}

			if v.Text != nil {
				fmt.Print(*v.Text)
			}
		},
	}
}
