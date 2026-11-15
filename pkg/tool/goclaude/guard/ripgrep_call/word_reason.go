package ripgrep_call

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"mvdan.cc/sh/v3/syntax"
	"slices"
	"strings"
)

func wordReason(w *syntax.Word) string {
	leading, readable := shell_script.LeadingCharacter(w)

	if !readable {
		return constant.RipgrepUnreadableReason
	}

	if leading != '-' {
		return ""
	}

	option, plain := shell_script.Argument(w)

	if !plain {
		return constant.RipgrepUnreadableReason
	}

	name, _, _ := strings.Cut(option, "=")

	if slices.Contains(constant.RipgrepProgramOptions, name) {
		return fmt.Sprintf(constant.RipgrepProgramReason, option)
	}

	return ""
}
