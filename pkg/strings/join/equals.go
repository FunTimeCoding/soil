package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Equals(s []string) string {
	return strings.Join(s, constant.Equals)
}
