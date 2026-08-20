package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func DoubleSpace(s []string) string {
	return strings.Join(s, constant.DoubleSpace)
}
