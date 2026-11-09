package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func deleteBackup(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var instance string
	result := &cobra.Command{
		Use:   "delete-backup [backup]",
		Short: "Delete a backup and its checksum",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitDone(
				(*c).DeleteBackup(
					context.Background(),
					arguments[0],
					&client.DeleteBackupParams{
						Instance: instancePointer(instance),
					},
				),
			)
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
