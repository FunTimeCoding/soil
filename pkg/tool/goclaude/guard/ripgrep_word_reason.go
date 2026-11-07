package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"slices"
	"strings"
)

func ripgrepWordReason(w *syntax.Word) string {
	leading, readable := leadingCharacter(w)

	if !readable {
		return constant.RipgrepUnreadableReason
	}

	if leading != '-' {
		return ""
	}

	option, plain := plainWord(w)

	if !plain {
		return constant.RipgrepUnreadableReason
	}

	name, _, _ := strings.Cut(option, "=")

	if slices.Contains(constant.RipgrepProgramOptions, name) {
		return fmt.Sprintf(constant.RipgrepProgramReason, option)
	}

	return ""
}
