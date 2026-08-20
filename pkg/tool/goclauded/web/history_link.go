package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func historyLink(
	page int,
	kinds []string,
) string {
	var parts []string

	for _, k := range kinds {
		parts = append(parts, key_value.Equals(constant.Kind, k))
	}

	if page > 1 {
		parts = append(parts, fmt.Sprintf("page=%d", page))
	}

	if len(parts) == 0 {
		return "/history"
	}

	return fmt.Sprintf("/history?%s", join.Ampersand(parts))
}
