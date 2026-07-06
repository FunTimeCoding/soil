package goclaude

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionBranch(c *command_context.Context) *cobra.Command {
	result := &cobra.Command{Use: "session"}
	result.AddCommand(sessionList(c))
	result.AddCommand(sessionShow(c))
	result.AddCommand(sessionPrint(c))
	result.AddCommand(sessionPeek(c))
	result.AddCommand(sessionEdit(c))
	result.AddCommand(sessionDelete(c))
	result.AddCommand(sessionExport(c))
	result.AddCommand(sessionTools(c))
	result.AddCommand(sessionHeatmap(c))
	result.AddCommand(sessionBashDump(c))
	result.AddCommand(sessionContext(c))
	result.AddCommand(sessionFind(c))
	result.AddCommand(sessionSweep(c))
	result.AddCommand(sessionBackfill(c))

	return result
}
