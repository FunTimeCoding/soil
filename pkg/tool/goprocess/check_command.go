package goprocess

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"github.com/spf13/cobra"
	"sort"
)

func checkCommand(t *terminal.Terminal) *cobra.Command {
	return &cobra.Command{
		Use:   "check",
		Short: "Validate Procfile and show entries",
		Run: func(
			m *cobra.Command,
			_ []string,
		) {
			procfilePath, e := m.Flags().GetString("file")
			errors.PanicOnError(e)
			entries, f := procfile.Parse(procfilePath)

			if f != nil {
				t.EmitDone(f)

				return
			}

			names := make([]string, len(entries))

			for i, entry := range entries {
				names[i] = entry.Name
			}

			sort.Strings(names)
			t.EmitValue(names, nil)
		},
	}
}
