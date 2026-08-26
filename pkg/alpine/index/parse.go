package index

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
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

		if v, okay := strings.CutPrefix(line, constant.IndexName); okay {
			current.Name = v
		}

		if v, okay := strings.CutPrefix(line, constant.IndexVersion); okay {
			current.Version = v
		}

		if v, okay := strings.CutPrefix(
			line,
			constant.IndexArchitecture,
		); okay {
			current.Architecture = v
		}
	}

	if current.Name != "" {
		result = append(result, current)
	}

	return result
}
