package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
)

func bashGroupKey(command string) string {
	command = strings.TrimSpace(command)

	if command == "" {
		return ""
	}

	parts := strings.Fields(command)

	if len(parts) == 0 {
		return ""
	}

	first := parts[0]

	if constant.BashNoise[first] {
		return ""
	}

	if strings.HasPrefix(first, "//") ||
		strings.HasPrefix(first, "{") ||
		strings.HasPrefix(first, "(") ||
		strings.HasPrefix(first, ")") ||
		strings.HasPrefix(first, "}") ||
		strings.HasPrefix(first, "\"") ||
		strings.HasPrefix(first, "'") ||
		strings.HasPrefix(first, "s/") ||
		strings.Contains(first, "=") {
		return ""
	}

	if constant.MultiWordPrefixes[first] && len(parts) > 1 {
		return fmt.Sprintf("%s %s", first, parts[1])
	}

	return first
}
