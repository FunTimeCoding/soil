package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"testing"
)

func TestExtractReplaces(t *testing.T) {
	assert.String(
		t,
		"replace (\n\tk8s.io/api => k8s.io/api v0.31.0\n)",
		go_mod.ExtractReplaces(
			`module github.com/example/project/v2

go 1.22.0

require (
	golang.org/x/crypto v0.31.0
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f
)

replace (
	k8s.io/api => k8s.io/api v0.31.0
)`,
		),
	)
}

func TestIsDeadTag(t *testing.T) {
	assert.True(
		t,
		go_mod.IsDeadTag(
			"go: github.com/funtimecoding/soil@v0.10.307: reading github.com/funtimecoding/soil/go.mod at revision v0.10.307: unknown revision v0.10.307",
		),
	)
}

func TestIsDeadTagNegative(t *testing.T) {
	assert.False(t, go_mod.IsDeadTag("go: module not found"))
}

func TestParseDeadTag(t *testing.T) {
	mod, version := go_mod.ParseDeadTag(
		"go: github.com/funtimecoding/soil@v0.10.307: reading github.com/funtimecoding/soil/go.mod at revision v0.10.307: unknown revision v0.10.307",
	)
	assert.String(t, "github.com/funtimecoding/soil", mod)
	assert.String(t, "v0.10.307", version)
}

func TestParseDeadTagNoMatch(t *testing.T) {
	mod, version := go_mod.ParseDeadTag("go: module not found")
	assert.String(t, "", mod)
	assert.String(t, "", version)
}

func TestReplaceReplaces(t *testing.T) {
	assert.String(
		t,
		"replace (\nb\n)\n",
		go_mod.ReplaceReplaces("replace (\na\n)\n", "replace (\nb\n)\n"),
	)
}
