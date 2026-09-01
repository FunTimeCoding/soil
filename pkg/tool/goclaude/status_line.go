package goclaude

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func statusLine(c *command_context.Context) *cobra.Command {
	result := &cobra.Command{
		Use:   "status-line",
		Short: "Render the status line and report context usage to goclauded",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			body, e := io.ReadAll(os.Stdin)
			errors.PanicOnError(e)
			errors.LogOnError(
				os.WriteFile(constant.StatusLineDumpFile, body, 0o644),
			)
			console.Line(RunStatusLine(c.Client(), body))
		},
	}

	return result
}
