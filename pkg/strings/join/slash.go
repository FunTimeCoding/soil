package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Slash(s []string) string {
	return strings.Join(s, constant.Slash)
}
