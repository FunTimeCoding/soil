package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func statusline(c *command_context.Context) *cobra.Command {
	result := &cobra.Command{
		Use:   "statusline",
		Short: "Render the status line and report context usage to goclauded",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			body, e := io.ReadAll(os.Stdin)
			errors.PanicOnError(e)
			errors.LogOnError(
				os.WriteFile(constant.StatuslineDumpFile, body, 0o644),
			)
			fmt.Println(RunStatusline(c.Client(), body))
		},
	}

	return result
}
