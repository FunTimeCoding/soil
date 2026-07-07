package pointer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestRoots(t *testing.T) {
	assert.Any(
		t,
		[]string{".claude", "doc", "pkg"},
		Roots([]string{
			"CLAUDE.md",
			"doc/ai/spec/naming.md",
			"doc/ai/runbook/lint.md",
			"pkg/lint/check.go",
			".claude/skills/sign-firefox/SKILL.md",
		}),
	)
	assert.Any(t, []string(nil), Roots([]string{"README.md"}))
	assert.Any(t, []string(nil), Roots(nil))
}
