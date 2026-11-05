package goclaude

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"github.com/spf13/cobra"
)

func New(
	c *command_context.Context,
	r face.Reporter,
) *cobra.Command {
	result := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	result.AddCommand(sessionBranch(c))
	result.AddCommand(messageBranch(c))
	result.AddCommand(register(c))
	result.AddCommand(sessionEnd(c))
	result.AddCommand(turnEnd(c))
	result.AddCommand(status(c))
	result.AddCommand(check(c))
	result.AddCommand(statusLine(c))
	result.AddCommand(guard.New(c.Terminal()))
	result.AddCommand(serveChannel(c, r))
	result.AddCommand(usage(c))
	result.AddCommand(cost(c))

	return result
}
