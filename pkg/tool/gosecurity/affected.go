package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/spf13/cobra"
)

func affected(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "affected <identifier>",
		Short: "Show the systems a CVE hits, with each one's versions and fix kind",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitValue(c.GetAffectedSystems(context.Background(), a[0]))
		},
	}
}
