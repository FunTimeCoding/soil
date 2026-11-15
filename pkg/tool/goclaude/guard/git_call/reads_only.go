package git_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"slices"
)

func (c *Call) readsOnly() bool {
	switch c.subcommand {
	case constant.GitStashSubcommand:
		return slices.Contains(constant.GitStashReadVerbs, c.stashVerb())
	case constant.GitMoveSubcommand:
		return c.dryRun()
	default:
		return false
	}
}
