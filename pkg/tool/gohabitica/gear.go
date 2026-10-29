package gohabitica

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
	"github.com/spf13/cobra"
)

func gear(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "gear",
		Short: "List owned gear",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.GetGear(context.Background()))
		},
	}
}
