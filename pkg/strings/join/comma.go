package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Comma(s []string) string {
	return strings.Join(s, constant.Comma)
}
