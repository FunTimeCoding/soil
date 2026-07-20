package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestClassify(t *testing.T) {
	roots := []string{".claude", ".claude-plugin", "doc", "pkg", "skills"}
	assert.String(
		t,
		pointer.Repository,
		pointer.Classify(
			"doc/ai/spec/naming.md",
			roots,
		),
	)
	assert.String(
		t,
		pointer.Repository,
		pointer.Classify(
			".claude/skills/sign-firefox/SKILL.md",
			roots,
		),
	)
	assert.String(
		t,
		pointer.Repository,
		pointer.Classify(".claude-plugin/plugin.json", roots),
	)
	assert.String(
		t,
		pointer.Repository,
		pointer.Classify("./doc/ai/spec", roots),
	)
	assert.String(
		t,
		pointer.Repository,
		pointer.Classify(
			"${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md",
			roots,
		),
	)
	assert.String(
		t,
		pointer.Unknown,
		pointer.Classify("tmp/gosec.json", roots),
	)
	assert.String(
		t,
		pointer.Unknown,
		pointer.Classify(constant.SoilModule, roots),
	)
	assert.String(
		t,
		pointer.Unknown,
		pointer.Classify("/chart-sessions", roots),
	)
	assert.String(t, pointer.Unknown, pointer.Classify("/api/goals", roots))
	assert.String(
		t,
		pointer.Unknown,
		pointer.Classify("/debug/pprof/", roots),
	)
	assert.String(t, pointer.Unknown, pointer.Classify("//nolint", roots))
	assert.String(t, pointer.Unknown, pointer.Classify("/etc/hosts", roots))
	assert.String(
		t,
		pointer.Unknown,
		pointer.Classify(
			"pkg/web/RecoveryMiddleware",
			roots,
		),
	)
	assert.String(
		t,
		pointer.Unknown,
		pointer.Classify(
			"pkg/provision/salt.Client",
			roots,
		),
	)
	assert.String(
		t,
		pointer.Unknown,
		pointer.Classify("pkg/check/memory.LocalLines()", roots),
	)
	assert.String(
		t,
		pointer.Placeholder,
		pointer.Classify(
			"doc/ai/runbook/<name>.md",
			roots,
		),
	)
	assert.String(
		t,
		pointer.Placeholder,
		pointer.Classify("pkg/tool/*.go", roots),
	)
	assert.String(
		t,
		pointer.Placeholder,
		pointer.Classify("$HOME/notes.md", roots),
	)
	assert.String(
		t,
		pointer.Placeholder,
		pointer.Classify(
			"${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md",
			roots,
		),
	)
	assert.String(
		t,
		pointer.Sibling,
		pointer.Classify("../github/soil/doc/ai/spec/naming.md", roots),
	)
	assert.String(
		t,
		pointer.Absolute,
		pointer.Classify(
			"/Users/example/notes.md",
			roots,
		),
	)
	assert.String(
		t,
		pointer.Locator,
		pointer.Classify("https://code.claude.com/docs/en/skills", roots),
	)
}
