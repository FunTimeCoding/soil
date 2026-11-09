package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func restoreBackup(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var instance string
	var target string
	result := &cobra.Command{
		Use:   "restore-backup [backup]",
		Short: "Restore a backup into a new database, or in place when the target is the instance's own",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			b := client.RestoreBackupJSONRequestBody{
				Instance: instancePointer(instance),
			}

			if target != "" {
				b.Target = &target
			}

			r, e := (*c).RestoreBackup(context.Background(), arguments[0], b)
			t.EmitValue(r, e)
		},
	}
	result.Flags().StringVar(
		&instance,
		"instance",
		"",
		"Instance name (optional when only one instance is configured)",
	)
	result.Flags().StringVar(
		&target,
		"target",
		"",
		"Database to restore into (default: a new <database>_restore_<backup>)",
	)

	return result
}
