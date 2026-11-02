package host

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func HasDot(s string) bool {
	return strings.Count(s, stringsConstant.Dot) > 0
}
