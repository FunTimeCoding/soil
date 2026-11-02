package service

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func ParseSearch(output string) map[string]string {
	result := make(map[string]string)

	for _, line := range strings.Split(output, stringsConstant.Unix) {
		name, path, okay := strings.Cut(line, stringsConstant.Colon)

		if !okay {
			continue
		}

		path = strings.TrimSpace(path)

		if !strings.HasPrefix(path, stringsConstant.Slash) ||
			strings.Contains(path, stringsConstant.Space) ||
			strings.Contains(name, stringsConstant.Space) {
			continue
		}

		result[path] = name
	}

	return result
}
