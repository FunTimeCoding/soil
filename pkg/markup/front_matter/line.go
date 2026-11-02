package front_matter

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (f *Front) Line(key string) int {
	prefix := join.Empty(key, stringsConstant.Colon)

	for i, line := range strings.Split(f.Raw, stringsConstant.Unix) {
		if strings.HasPrefix(line, prefix) {
			return i + 2
		}
	}

	return 0
}
