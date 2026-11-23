package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/spf13/cobra"
)

func systems(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "systems",
		Short: "List the selected systems by priority with coverage, finding counts, top score and kernel state",
		Args:  cobra.NoArgs,
		Run: func(
			o *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.ListSystems(context.Background(), selection(o)))
		},
	}
	selectionFlags(result)

	return result
}
