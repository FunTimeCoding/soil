package host

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func HasDot(s string) bool {
	return strings.Count(s, constant.Dot) > 0
}
