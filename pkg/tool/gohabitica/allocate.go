package gohabitica

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
	"github.com/spf13/cobra"
)

func allocate(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "allocate <stat>",
		Short: "Allocate a stat point: str, con, int, per",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(c.AllocateStat(context.Background(), arguments[0]))
		},
	}
}
