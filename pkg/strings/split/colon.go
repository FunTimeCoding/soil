package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Colon(s string) []string {
	return strings.Split(s, constant.Colon)
}
