package goalertlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/face"
	"github.com/spf13/cobra"
)

func status(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the record count, latest identifier and last poll",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.GetStatus(context.Background()))
		},
	}
}
