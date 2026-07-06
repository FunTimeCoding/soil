package go_mod

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestIsDeadTag(t *testing.T) {
	assert.True(
		t,
		isDeadTag(
			"go: github.com/funtimecoding/soil@v0.10.307: reading github.com/funtimecoding/soil/go.mod at revision v0.10.307: unknown revision v0.10.307",
		),
	)
}

func TestIsDeadTagNegative(t *testing.T) {
	assert.False(
		t,
		isDeadTag("go: module not found"),
	)
}

func TestParseDeadTag(t *testing.T) {
	mod, version := parseDeadTag(
		"go: github.com/funtimecoding/soil@v0.10.307: reading github.com/funtimecoding/soil/go.mod at revision v0.10.307: unknown revision v0.10.307",
	)
	assert.String(t, "github.com/funtimecoding/soil", mod)
	assert.String(t, "v0.10.307", version)
}

func TestParseDeadTagNoMatch(t *testing.T) {
	mod, version := parseDeadTag("go: module not found")
	assert.String(t, "", mod)
	assert.String(t, "", version)
}
