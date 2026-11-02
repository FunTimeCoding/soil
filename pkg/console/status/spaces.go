package status

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func spaces(indent int) string {
	return strings.Repeat(stringsConstant.DoubleSpace, indent)
}
