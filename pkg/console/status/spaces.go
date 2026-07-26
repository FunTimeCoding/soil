package status

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func spaces(indent int) string {
	return strings.Repeat(constant.DoubleSpace, indent)
}
