package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/grant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/python_replacement"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
)

func GrantedBlockReason(
	system string,
	command string,
	g *grant.Grant,
) string {
	s := shell_script.Parse(command)

	if !s.Parsed && constant.GitInvocation.MatchString(command) {
		return fmt.Sprintf(constant.GitUnreadableReason, "command")
	}

	if b := gitBlockReason(s, command, g); b != "" {
		return b
	}

	unparsed := s.Unparsed()

	if unparsed && constant.NpxInvocation.MatchString(command) ||
		s.Any(npxCall) {
		return constant.NpxReason
	}

	if unparsed && constant.PipInvocation.MatchString(command) ||
		s.Any(pipInstallCall) {
		return constant.PipReason
	}

	if python_replacement.Unrestored(s) {
		return constant.SingleReplacementReason
	}

	if system != "darwin" {
		return ""
	}

	if unparsed && constant.SedInvocation.MatchString(command) ||
		s.Any(sedInPlaceCall) {
		return constant.SedReason
	}

	if unparsed && constant.XargsInvocation.MatchString(command) ||
		s.Any(xargsUnsupportedFlagCall) {
		return constant.XargsReason
	}

	return ""
}
