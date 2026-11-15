package base

import (
	"os"
	"strings"
	"testing"
)

func LaunchCount(
	t *testing.T,
	marker string,
) int {
	t.Helper()
	content, e := os.ReadFile(marker)

	if e != nil {
		return 0
	}

	return len(strings.Fields(string(content)))
}
