package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Question(s ...string) string {
	return strings.Join(s, constant.Question)
}
