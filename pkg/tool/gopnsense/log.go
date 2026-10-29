package gopnsense

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/spf13/cobra"
)

func log(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var limit int
	result := &cobra.Command{
		Use:   "log",
		Short: "Read recent firewall log records",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.Log(context.Background(), limit))
		},
	}
	result.Flags().IntVar(&limit, "limit", 0, "maximum number of records")

	return result
}
