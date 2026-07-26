package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Underscore(s []string) string {
	return strings.Join(s, constant.Underscore)
}
