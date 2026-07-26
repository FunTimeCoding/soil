package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Dot(s []string) string {
	return strings.Join(s, constant.Dot)
}
