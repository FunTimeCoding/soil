package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func parseFixtureRule(line string) (string, bool) {
	trimmed := strings.TrimSpace(line)

	if !strings.HasPrefix(trimmed, constant.FixturePrefix) {
		return "", false
	}

	return strings.TrimSpace(
		strings.TrimPrefix(trimmed, constant.FixturePrefix),
	), true
}
