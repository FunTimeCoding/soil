package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestRoots(t *testing.T) {
	assert.Any(
		t,
		[]string{".claude", "doc", "pkg"},
		pointer.Roots(
			[]string{
				"CLAUDE.md",
				"doc/ai/spec/naming.md",
				"doc/ai/runbook/lint.md",
				"pkg/lint/check.go",
				".claude/skills/sign-firefox/SKILL.md",
			},
		),
	)
	assert.Any(t, []string(nil), pointer.Roots([]string{"README.md"}))
	assert.Any(t, []string(nil), pointer.Roots(nil))
}
