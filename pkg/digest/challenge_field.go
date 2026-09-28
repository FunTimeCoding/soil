package digest

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func challengeField(
	challenge string,
	name string,
) string {
	marker := join.Empty(name, `="`)
	start := strings.Index(challenge, marker)

	if start < 0 {
		return ""
	}

	rest := challenge[start+len(marker):]
	end := strings.Index(rest, `"`)

	if end < 0 {
		return ""
	}

	return rest[:end]
}
