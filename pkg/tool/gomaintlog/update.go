package gomaintlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/face"
	"github.com/spf13/cobra"
	"strconv"
)

func update(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "update <id>",
		Short: "Replace a maintenance entry's fields",
		Args:  cobra.ExactArgs(1),
		Run: func(
			o *cobra.Command,
			a []string,
		) {
			identifier, e := strconv.Atoi(a[0])
			cobra.CheckErr(e)
			t.EmitValue(
				c.UpdateEntry(context.Background(), identifier, request(o)),
			)
		},
	}
	entryFlags(result)

	return result
}
