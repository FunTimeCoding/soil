package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func documentsLink(
	basePath string,
	page int,
	sourceType string,
) string {
	var parts []string

	if sourceType != "" {
		parts = append(
			parts,
			fmt.Sprintf("%s=%s", constant.SourceType, sourceType),
		)
	}

	if page > 1 {
		parts = append(parts, fmt.Sprintf("page=%d", page))
	}

	if len(parts) == 0 {
		return basePath
	}

	return fmt.Sprintf("%s?%s", basePath, strings.Join(parts, "&"))
}
