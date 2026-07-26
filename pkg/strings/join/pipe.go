package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Pipe(s []string) string {
	return strings.Join(s, constant.Pipe)
}
