package probe

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"strings"
)

func candidates(lines []string) map[string]string {
	result := map[string]string{}

	for _, line := range lines {
		name, rest, okay := strings.Cut(line, constant.SuiteSeparator)

		if !okay {
			continue
		}

		f := strings.Fields(rest)

		if len(f) < 3 {
			continue
		}

		if _, seen := result[name]; !seen {
			result[name] = f[1]
		}

		result[join.Colon(name, f[2])] = f[1]
	}

	return result
}
