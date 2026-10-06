package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Message(
	verdict constant.Verdict,
	hint string,
) string {
	if hint == "" {
		return constant.VerdictTexts[verdict]
	}

	return join.Empty(
		constant.VerdictTexts[verdict],
		constant.HintSeparator,
		hint,
	)
}
