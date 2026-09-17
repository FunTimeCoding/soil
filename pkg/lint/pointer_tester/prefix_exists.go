package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func prefixExists(all []string) func(string, string) bool {
	return func(
		directory string,
		prefix string,
	) bool {
		full := join.Empty(directory, constant.Slash, prefix)

		for _, p := range all {
			if strings.HasPrefix(p, full) {
				return true
			}
		}

		return false
	}
}
