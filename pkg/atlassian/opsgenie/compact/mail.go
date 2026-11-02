package compact

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"strings"
)

func Mail(s string) string {
	if strings.Contains(s, stringsConstant.At) {
		first, _ := key_value.At(s)

		return first
	}

	return s
}
