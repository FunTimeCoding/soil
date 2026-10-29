package gomaintlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/face"
	"github.com/spf13/cobra"
)

func add(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "add",
		Short: "Record a maintenance entry",
		Args:  cobra.NoArgs,
		Run: func(
			o *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.PostEntry(context.Background(), request(o)))
		},
	}
	entryFlags(result)

	return result
}
