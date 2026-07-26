package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Space(s string) []string {
	return strings.Split(s, constant.Space)
}
