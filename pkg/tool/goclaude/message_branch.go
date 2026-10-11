package goclaude

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func messageBranch(c *command_context.Context) *cobra.Command {
	result := &cobra.Command{Use: "message"}
	result.AddCommand(messageRead(c))

	return result
}
