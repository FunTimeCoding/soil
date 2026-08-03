package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"path/filepath"
	"strings"
)

func ExtractTitle(
	content string,
	filename string,
) string {
	match := constant.HeadingPattern.FindStringSubmatch(content)

	if match != nil {
		return strings.TrimSpace(match[1])
	}

	base := filepath.Base(filename)
	extension := filepath.Ext(base)

	return strings.TrimSuffix(base, extension)
}
