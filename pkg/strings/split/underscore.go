package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Underscore(s string) []string {
	return strings.Split(s, constant.Underscore)
}
