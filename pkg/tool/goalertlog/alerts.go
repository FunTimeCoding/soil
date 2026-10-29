package goalertlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/spf13/cobra"
)

func alerts(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "alerts <name>",
		Short: "List every recorded alert with this name",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitValue(
				c.GetAlerts(
					context.Background(),
					&client.GetAlertsParams{Name: a[0]},
				),
			)
		},
	}
}
