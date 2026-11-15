package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Message(
	key string,
	hint string,
) string {
	if hint == "" {
		return constant.PointerTexts[key]
	}

	return join.Empty(constant.PointerTexts[key], constant.HintSeparator, hint)
}
