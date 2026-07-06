package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
)

func historyLink(
	page int,
	kinds []string,
) string {
	var parts []string

	for _, k := range kinds {
		parts = append(
			parts,
			fmt.Sprintf("%s=%s", constant.Kind, k),
		)
	}

	if page > 1 {
		parts = append(parts, fmt.Sprintf("page=%d", page))
	}

	if len(parts) == 0 {
		return "/history"
	}

	return fmt.Sprintf("/history?%s", strings.Join(parts, "&"))
}
