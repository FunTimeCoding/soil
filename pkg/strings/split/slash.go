package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Slash(s string) []string {
	return strings.Split(s, constant.Slash)
}
