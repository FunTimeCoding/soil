package index

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
)

func parse(content string) []*Entry {
	var result []*Entry
	current := &Entry{}

	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			if current.Name != "" {
				result = append(result, current)
			}

			current = &Entry{}

			continue
		}

		if v, okay := strings.CutPrefix(line, constant.AlpineIndexName); okay {
			current.Name = v
		}

		if v, okay := strings.CutPrefix(line, constant.AlpineIndexVersion); okay {
			current.Version = v
		}

		if v, okay := strings.CutPrefix(
			line,
			constant.AlpineIndexArchitecture,
		); okay {
			current.Architecture = v
		}
	}

	if current.Name != "" {
		result = append(result, current)
	}

	return result
}
