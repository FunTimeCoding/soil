package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"path"
)

func ripgrepReason(
	call *syntax.CallExpr,
	configured bool,
) string {
	words, fromInput := wrappedCommand(call.Args)

	if len(words) == 0 {
		return ""
	}

	if name, _ := literalWord(words[0]); path.Base(name) != constant.RipgrepCommand {
		return ""
	}

	if fromInput || configured {
		return constant.RipgrepUnreadableReason
	}

	for _, w := range words[1:] {
		if w.Lit() == "--" {
			return ""
		}

		if r := ripgrepWordReason(w); r != "" {
			return r
		}
	}

	return ""
}
