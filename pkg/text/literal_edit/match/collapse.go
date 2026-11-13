package match

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func collapse(s string) string {
	return strings.Join(strings.Fields(s), stringsConstant.Space)
}
