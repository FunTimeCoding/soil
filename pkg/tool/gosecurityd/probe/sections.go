package probe

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"strings"
)

func sections(output string) map[string][]string {
	result := map[string][]string{}
	current := ""

	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimRight(line, "\r")

		if strings.HasPrefix(line, constant.SectionPrefix) {
			current = line
			result[current] = []string{}

			continue
		}

		if current != "" && strings.TrimSpace(line) != "" {
			result[current] = append(result[current], line)
		}
	}

	return result
}
