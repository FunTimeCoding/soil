package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Comma(s string) []string {
	return strings.Split(s, constant.Comma)
}
