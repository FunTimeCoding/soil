package key_value

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Space(s string) (string, string) {
	p := strings.SplitN(s, constant.Space, 2)

	switch len(p) {
	case 1:
		return p[0], ""
	case 2:
		return p[0], p[1]
	}

	return "", ""
}
