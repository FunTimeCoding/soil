package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/git_call"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/python_replacement"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
)

func BlockReason(
	system string,
	command string,
) string {
	s := shell_script.Parse(command)

	if !s.Parsed && constant.GitInvocation.MatchString(command) {
		return fmt.Sprintf(constant.GitUnreadableReason, "command")
	}

	for _, c := range git_call.Calls(s) {
		if b := c.BlockReason(); b != "" {
			return b
		}
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
