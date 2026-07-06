package goprocess

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"github.com/spf13/cobra"
	"sort"
)

func checkCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "check",
		Short: "Validate Procfile and show entries",
		RunE: func(
			cmd *cobra.Command,
			_ []string,
		) error {
			cmd.SilenceUsage = true
			procfilePath, e := cmd.Flags().GetString("file")
			errors.PanicOnError(e)
			entries, e := procfile.Parse(procfilePath)

			if e != nil {
				return e
			}

			names := make([]string, len(entries))

			for i, entry := range entries {
				names[i] = entry.Name
			}

			sort.Strings(names)
			fmt.Printf(
				"valid procfile detected (%s)\n",
				join.CommaSpace(names),
			)

			return nil
		},
	}
}
