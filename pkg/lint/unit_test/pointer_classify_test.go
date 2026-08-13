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
		"repository",
		pointer.Classify("doc/ai/spec/naming.md", roots),
	)
	assert.String(
		t,
		"repository",
		pointer.Classify(".claude/skills/sign-firefox/SKILL.md", roots),
	)
	assert.String(
		t,
		"repository",
		pointer.Classify(".claude-plugin/plugin.json", roots),
	)
	assert.String(t, "repository", pointer.Classify("./doc/ai/spec", roots))
	assert.String(
		t,
		"repository",
		pointer.Classify(
			"${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/constant.md",
			roots,
		),
	)
	assert.String(t, "unknown", pointer.Classify("tmp/gosec.json", roots))
	assert.String(t, "unknown", pointer.Classify(constant.SoilModule, roots))
	assert.String(t, "unknown", pointer.Classify("/chart-sessions", roots))
	assert.String(t, "unknown", pointer.Classify("/api/goals", roots))
	assert.String(t, "unknown", pointer.Classify("/debug/pprof/", roots))
	assert.String(t, "unknown", pointer.Classify("//nolint", roots))
	assert.String(t, "unknown", pointer.Classify("/etc/hosts", roots))
	assert.String(
		t,
		"unknown",
		pointer.Classify("pkg/web/RecoveryMiddleware", roots),
	)
	assert.String(
		t,
		"unknown",
		pointer.Classify("pkg/provision/salt.Client", roots),
	)
	assert.String(
		t,
		"unknown",
		pointer.Classify("pkg/check/memory.LocalLines()", roots),
	)
	assert.String(
		t,
		"placeholder",
		pointer.Classify("doc/ai/runbook/<name>.md", roots),
	)
	assert.String(t, "placeholder", pointer.Classify("pkg/tool/*.go", roots))
	assert.String(t, "placeholder", pointer.Classify("$HOME/notes.md", roots))
	assert.String(
		t,
		"placeholder",
		pointer.Classify("${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md", roots),
	)
	assert.String(
		t,
		"sibling",
		pointer.Classify("../github/soil/doc/ai/spec/naming.md", roots),
	)
	assert.String(
		t,
		"absolute",
		pointer.Classify("/Users/example/notes.md", roots),
	)
	assert.String(
		t,
		"locator",
		pointer.Classify("https://code.claude.com/docs/en/skills", roots),
	)
}
