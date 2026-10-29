package goalertlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/spf13/cobra"
)

func alertsByIdentifier(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var identifiers []int
	result := &cobra.Command{
		Use:   "alerts-by-identifier",
		Short: "Read recorded alerts by identifier",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(
				c.GetAlertsByIdentifier(
					context.Background(),
					&client.GetAlertsByIdentifierParams{Identifier: identifiers},
				),
			)
		},
	}
	result.Flags().IntSliceVar(
		&identifiers,
		"identifier",
		nil,
		"identifier to read, repeatable",
	)

	return result
}
