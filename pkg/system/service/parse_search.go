package service

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func ParseSearch(output string) map[string]string {
	result := make(map[string]string)

	for _, line := range strings.Split(output, constant.Unix) {
		name, path, okay := strings.Cut(line, constant.Colon)

		if !okay {
			continue
		}

		path = strings.TrimSpace(path)

		if !strings.HasPrefix(path, constant.Slash) ||
			strings.Contains(path, constant.Space) ||
			strings.Contains(name, constant.Space) {
			continue
		}

		result[path] = name
	}

	return result
}
