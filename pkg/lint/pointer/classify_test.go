package pointer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"testing"
)

func TestClassify(t *testing.T) {
	roots := []string{".claude", ".claude-plugin", "doc", "pkg", "skills"}
	assert.String(t, Repository, Classify("doc/ai/spec/naming.md", roots))
	assert.String(
		t,
		Repository,
		Classify(
			".claude/skills/sign-firefox/SKILL.md",
			roots,
		),
	)
	assert.String(t, Repository, Classify(".claude-plugin/plugin.json", roots))
	assert.String(t, Repository, Classify("./doc/ai/spec", roots))
	assert.String(
		t,
		Repository,
		Classify("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md", roots),
	)
	assert.String(t, Unknown, Classify("tmp/gosec.json", roots))
	assert.String(t, Unknown, Classify(constant.SoilModule, roots))
	assert.String(t, Unknown, Classify("/chart-sessions", roots))
	assert.String(t, Unknown, Classify("/api/goals", roots))
	assert.String(t, Unknown, Classify("/debug/pprof/", roots))
	assert.String(t, Unknown, Classify("//nolint", roots))
	assert.String(t, Unknown, Classify("/etc/hosts", roots))
	assert.String(t, Unknown, Classify("pkg/web/RecoveryMiddleware", roots))
	assert.String(t, Placeholder, Classify("doc/ai/runbook/<name>.md", roots))
	assert.String(t, Placeholder, Classify("pkg/tool/*.go", roots))
	assert.String(t, Placeholder, Classify("$HOME/notes.md", roots))
	assert.String(
		t,
		Placeholder,
		Classify("${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md", roots),
	)
	assert.String(
		t,
		Sibling,
		Classify("../github/soil/doc/ai/spec/naming.md", roots),
	)
	assert.String(t, Absolute, Classify("/Users/example/notes.md", roots))
	assert.String(
		t,
		Locator,
		Classify("https://code.claude.com/docs/en/skills", roots),
	)
}
