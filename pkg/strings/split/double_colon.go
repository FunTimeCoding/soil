package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func DoubleColon(s string) []string {
	return strings.Split(s, constant.DoubleColon)
}
