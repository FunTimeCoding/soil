package goalertlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/spf13/cobra"
)

func alertsAfter(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var identifier int
	var limit int
	result := &cobra.Command{
		Use:   "alerts-after",
		Short: "Page through recorded alerts after an identifier",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			p := &client.GetAlertsAfterParams{Identifier: identifier}

			if limit > 0 {
				p.Limit = &limit
			}

			t.EmitValue(c.GetAlertsAfter(context.Background(), p))
		},
	}
	result.Flags().IntVar(
		&identifier,
		"identifier",
		0,
		"return alerts after this identifier",
	)
	result.Flags().IntVar(&limit, "limit", 0, "records per page, at most 25")

	return result
}
