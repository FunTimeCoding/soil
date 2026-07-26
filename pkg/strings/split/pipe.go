package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Pipe(s string) []string {
	return strings.Split(s, constant.Pipe)
}
