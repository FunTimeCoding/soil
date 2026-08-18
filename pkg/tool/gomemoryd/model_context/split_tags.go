package model_context

import (
	separator "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strings"
)

func splitTags(raw string) ([]string, bool) {
	var result []string
	stripped := false

	for _, one := range strings.Split(raw, separator.Comma) {
		clean := strings.Trim(one, constant.TagCutset)

		if clean != strings.TrimSpace(one) {
			stripped = true
		}

		if clean != "" {
			result = append(result, clean)
		}
	}

	return result, stripped
}
