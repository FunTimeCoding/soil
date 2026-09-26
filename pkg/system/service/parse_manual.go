package service

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func ParseManual(output string) map[string]bool {
	result := make(map[string]bool)

	for _, line := range strings.Split(output, constant.Unix) {
		if name := strings.TrimSpace(line); name != "" {
			result[name] = true
		}
	}

	return result
}
