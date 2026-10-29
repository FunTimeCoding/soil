package gomaintlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/face"
	"github.com/spf13/cobra"
	"strconv"
)

func remove(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a maintenance entry",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			identifier, e := strconv.Atoi(a[0])
			cobra.CheckErr(e)
			t.EmitDone(c.DeleteEntry(context.Background(), identifier))
		},
	}
}
