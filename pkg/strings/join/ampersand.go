package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Ampersand(s []string) string {
	return strings.Join(s, constant.Ampersand)
}
