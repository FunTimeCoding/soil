package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func Extract(line string) []string {
	var result []string
	parts := strings.Split(line, "`")

	for i := 1; i < len(parts)-1; i += 2 {
		if IsPath(parts[i]) {
			result = append(result, parts[i])
		}
	}

	for _, m := range constant.LinkTarget.FindAllStringSubmatch(line, -1) {
		if IsPath(m[1]) {
			result = append(result, m[1])
		}
	}

	return result
}
