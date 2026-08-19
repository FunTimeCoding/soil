package web

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func optionalList(
	target **[]string,
	value string,
) {
	if value == "" {
		return
	}

	var result []string

	for _, item := range strings.Split(value, constant.Comma) {
		if trimmed := strings.TrimSpace(item); trimmed != "" {
			result = append(result, trimmed)
		}
	}

	if len(result) == 0 {
		return
	}

	*target = &result
}
