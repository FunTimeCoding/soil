package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Dot(s string) []string {
	return strings.Split(s, constant.Dot)
}
