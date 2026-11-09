package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func createBackup(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var instance string
	result := &cobra.Command{
		Use:   "create-backup",
		Short: "Dump an instance's database to a new backup",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := (*c).CreateBackup(
				context.Background(),
				client.CreateBackupJSONRequestBody{
					Instance: instancePointer(instance),
				},
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
