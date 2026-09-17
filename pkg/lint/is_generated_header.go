package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func IsGeneratedHeader(content string) bool {
	for _, line := range strings.SplitN(content, "\n", 10) {
		trimmed := strings.TrimSpace(line)

		if !strings.HasPrefix(trimmed, constant.CommentPrefix) && trimmed != "" {
			return false
		}

		if strings.HasPrefix(trimmed, "// Code generated") &&
			strings.HasSuffix(trimmed, "DO NOT EDIT.") {
			return true
		}
	}

	return false
}
