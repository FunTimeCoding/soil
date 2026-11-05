package basic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"strings"
)

func reason(
	body []byte,
	fallback string,
) string {
	text := strings.TrimSpace(string(body))

	if text == "" || strings.ContainsAny(text[:1], constant.LokiMarkupOpenings) {
		return fallback
	}

	return text
}
