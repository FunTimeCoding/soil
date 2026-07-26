package compact

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"strings"
)

func Mail(s string) string {
	if strings.Contains(s, constant.At) {
		first, _ := key_value.At(s)

		return first
	}

	return s
}
