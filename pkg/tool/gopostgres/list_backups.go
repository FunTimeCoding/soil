package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func listBackups(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var instance string
	result := &cobra.Command{
		Use:   "list-backups",
		Short: "List an instance's backups, newest first",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := (*c).ListBackups(
				context.Background(),
				&client.ListBackupsParams{Instance: instancePointer(instance)},
			)
			t.EmitValue(r, e)
		},
	}
	result.Flags().StringVar(
		&instance,
		"instance",
		"",
		"Instance name (optional when only one instance is configured)",
	)

	return result
}
