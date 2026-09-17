package front_matter

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (f *Front) Line(key string) int {
	prefix := join.Empty(key, constant.Colon)

	for i, line := range strings.Split(f.Raw, constant.Unix) {
		if strings.HasPrefix(line, prefix) {
			return i + 2
		}
	}

	return 0
}
