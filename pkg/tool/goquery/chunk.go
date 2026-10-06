package goquery

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
	"path/filepath"
)

func chunk(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var snapshot, compare, restructure bool
	result := &cobra.Command{
		Use:   "chunk [file]",
		Short: "Diagnose how a file chunks; exit 1 if a chunk is too big or --restructure changed words",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			path := arguments[0]
			body := system.ReadFile(filepath.Dir(path), filepath.Base(path))
			failed := printPreview(c, filepath.Base(path), body) > 0

			if snapshot {
				system.SaveFile(snapshotPath(path, true), body)
			}

			if compare || restructure {
				changed := compareSnapshot(t, path, body)

				if restructure && changed > 0 {
					failed = true
				}
			}

			if failed {
				t.Exit(1)
			}
		},
	}
	result.Flags().BoolVar(
		&snapshot,
		constant.SnapshotFlag,
		false,
		"Store the file as it is now, to compare against after editing",
	)
	result.Flags().BoolVar(
		&compare,
		constant.CompareFlag,
		false,
		"Report words added and removed per section since the snapshot",
	)
	result.Flags().BoolVar(
		&restructure,
		constant.RestructureFlag,
		false,
		"Compare, and exit 1 if any body word was added or removed",
	)

	return result
}
