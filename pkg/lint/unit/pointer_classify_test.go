package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestClassify(t *testing.T) {
	roots := []string{".claude", ".claude-plugin", "doc", "pkg", "skills"}
	assert.String(t, "repository", classify("doc/ai/spec/naming.md", roots))
	assert.String(
		t,
		"repository",
		classify(".claude/skills/sign-firefox/SKILL.md", roots),
	)
	assert.String(
		t,
		"repository",
		classify(".claude-plugin/plugin.json", roots),
	)
	assert.String(t, "repository", classify("./doc/ai/spec", roots))
	assert.String(
		t,
		"repository",
		classify("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/constant.md", roots),
	)
	assert.String(t, "short", classify("tmp/gosec.json", roots))
	assert.String(t, "short", classify(constant.SoilModule, roots))
	assert.String(t, "command", classify("/chart-sessions", roots))
	assert.String(t, "command", classify("/soil:lint", roots))
	assert.String(t, "system", classify("/api/goals", roots))
	assert.String(t, "system", classify("/debug/pprof/", roots))
	assert.String(t, "pattern", classify("//nolint", roots))
	assert.String(t, "system", classify("/etc/hosts", roots))
	assert.String(t, "route", classify("route:/api/goals", roots))
	assert.String(t, "path", classify("path:/bin/true", roots))
	assert.String(t, "pattern", classify("s/a/b/g", roots))
	assert.String(t, "import", classify("\"example.org/module\"", roots))
	assert.String(
		t,
		"repository",
		classify("pkg/web/RecoveryMiddleware", roots),
	)
	assert.String(t, "repository", classify("pkg/provision/salt.Client", roots))
	assert.String(
		t,
		"repository",
		classify("pkg/check/memory.LocalLines()", roots),
	)
	assert.String(t, "symbol", classify("go:pkg/provision/salt.Client", roots))
	assert.String(
		t,
		"symbol",
		classify("go:../github/soil/pkg/provision/salt.Client", roots),
	)
	assert.String(t, "placeholder", classify("go:pkg/<name>/Symbol", roots))
	assert.String(t, "placeholder", classify("doc/ai/runbook/<name>.md", roots))
	assert.String(t, "placeholder", classify("pkg/tool/*.go", roots))
	assert.String(t, "placeholder", classify("$HOME/notes.md", roots))
	assert.String(
		t,
		"placeholder",
		classify("${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md", roots),
	)
	assert.String(
		t,
		"sibling",
		classify("../github/soil/doc/ai/spec/naming.md", roots),
	)
	assert.String(t, "absolute", classify("/Users/example/notes.md", roots))
	assert.String(
		t,
		"locator",
		classify("https://code.claude.com/docs/en/skills", roots),
	)
}

func classify(
	s string,
	roots []string,
) string {
	return string(pointer.Classify(s, roots))
}
