package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Colon(s ...string) string {
	return strings.Join(s, constant.Colon)
}
