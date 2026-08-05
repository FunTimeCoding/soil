package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func ExtractBareLinks(line string) []string {
	var result []string

	for _, m := range constant.LinkTarget.FindAllStringSubmatch(line, -1) {
		target := m[1]

		if IsPath(target) ||
			target == "" ||
			strings.ContainsAny(target, " \t") ||
			strings.Contains(target, "://") ||
			strings.HasPrefix(target, "#") ||
			strings.ContainsAny(target, "<>*$") ||
			!strings.ContainsRune(target, '.') {
			continue
		}

		result = append(result, target)
	}

	return result
}
