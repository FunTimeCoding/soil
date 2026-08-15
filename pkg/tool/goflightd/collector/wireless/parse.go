package wireless

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"strings"
)

func Parse(text string) map[string]string {
	result := make(map[string]string)
	section := ""

	for _, l := range strings.Split(text, "\n") {
		trimmed := strings.TrimSpace(l)

		if trimmed == "" || strings.Trim(trimmed, "—-=") == "" {
			continue
		}

		if !strings.Contains(trimmed, ":") {
			section = strings.ToLower(trimmed)

			continue
		}

		if section == "" {
			continue
		}

		parts := strings.SplitN(trimmed, ":", 2)
		key := strings.TrimSpace(parts[0])

		if key == "" {
			continue
		}

		result[join.Empty(section, constant.KeySeparator, key)] =
			strings.TrimSpace(parts[1])
	}

	return result
}
