package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Space(s ...string) string {
	return strings.Join(s, constant.Space)
}
