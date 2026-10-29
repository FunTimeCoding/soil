package gohabitica

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
	"github.com/spf13/cobra"
)

func statistic(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "statistic",
		Short: "Get user statistic",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.GetStats(context.Background()))
		},
	}
}
