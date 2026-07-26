package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func DoubleColon(s []string) string {
	return strings.Join(s, constant.DoubleColon)
}
