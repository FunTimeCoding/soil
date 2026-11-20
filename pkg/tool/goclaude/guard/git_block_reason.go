package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/git_call"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/grant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
)

func gitBlockReason(
	s *shell_script.Script,
	command string,
	g *grant.Grant,
) string {
	located := g != nil && locatable(s, command)

	for _, c := range git_call.Calls(s) {
		b := c.BlockReason()

		if b == "" {
			continue
		}

		if !located || !c.Grantable() {
			return b
		}

		repository := c.Repository(g.Directory())

		if repository != "" && g.Holds(repository) {
			g.Note(
				fmt.Sprintf(
					constant.GitGrantedNote,
					c.Subcommand(),
					constant.GitWriteLabel,
					repository,
				),
			)

			continue
		}

		if f := g.Failure(); f != nil {
			return fmt.Sprintf(
				constant.GitGrantUnreadReason,
				c.Subcommand(),
				constant.GitWriteLabel,
				f,
			)
		}

		return b
	}

	return ""
}
