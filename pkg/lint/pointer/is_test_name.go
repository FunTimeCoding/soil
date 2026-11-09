package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func IsTestName(s string) bool {
	name := strings.TrimSuffix(s, constant.CallSuffix)

	return constant.TestSpan.MatchString(name) &&
		!strings.Contains(name, constant.TestPlaceholder)
}
