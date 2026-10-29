package goraid

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/face"
	"github.com/spf13/cobra"
)

func reports(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "reports",
		Short: "List generated reports",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.GetReports(context.Background()))
		},
	}
}
