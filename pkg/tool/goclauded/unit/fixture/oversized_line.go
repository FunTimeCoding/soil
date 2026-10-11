package fixture

import (
	"strings"

	"github.com/funtimecoding/soil/pkg/strings/join"
)

func OversizedLine() string {
	return join.Empty(
		`{"type":"assistant","timestamp":"2026-08-16T00:30:00.000Z",`,
		`"message":{"role":"assistant","content":[{"type":"text","text":"`,
		strings.Repeat("x", 2*1024*1024),
		`"}]}}`,
	)
}
