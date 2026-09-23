package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func isDirective(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), constant.DirectivePrefix)
}
