package goalpine

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/face"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/client"
	"github.com/spf13/cobra"
)

func packages(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "packages [name]",
		Short: "List packages in the repository index",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			p := &client.GetPackagesParams{}

			if len(arguments) > 0 {
				p.Name = &arguments[0]
			}

			t.EmitValue(c.GetPackages(context.Background(), p))
		},
	}
}
