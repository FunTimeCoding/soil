package gosecret

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/spf13/cobra"
)

func listCommand(
	t *terminal.Terminal,
	c face.Client,
) *cobra.Command {
	return &cobra.Command{
		Use:   constant.ListUsage,
		Short: "List secrets and folders below a path; folders end in a slash",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			p := &client.ListSecretsParams{}

			if len(arguments) > 0 {
				p.Path = &arguments[0]
			}

			t.EmitValue(c.ListSecrets(context.Background(), p))
		},
	}
}
