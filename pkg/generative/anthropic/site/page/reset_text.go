package page

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"strings"
)

func resetText(s string) string {
	i := strings.Index(s, constant.UsageResetMarker)

	if i < 0 {
		return ""
	}

	return strings.TrimSpace(s[i+len(constant.UsageResetMarker):])
}
