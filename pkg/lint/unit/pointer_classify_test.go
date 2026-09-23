package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"testing"
)

func TestClassify(t *testing.T) {
	roots := []string{".claude", ".claude-plugin", "doc", "pkg", "skills"}
	assert.String(
		t,
		"repository",
		lint_tester.Classify("doc/ai/spec/naming.md", roots),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify(".claude/skills/sign-firefox/SKILL.md", roots),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify(".claude-plugin/plugin.json", roots),
	)
	assert.String(t, "repository", lint_tester.Classify("./doc/ai/spec", roots))
	assert.String(
		t,
		"repository",
		lint_tester.Classify(
			"${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/constant.md",
			roots,
		),
	)
	assert.String(t, "short", lint_tester.Classify("tmp/gosec.json", roots))
	assert.String(t, "short", lint_tester.Classify(constant.SoilModule, roots))
	assert.String(t, "command", lint_tester.Classify("/chart-sessions", roots))
	assert.String(t, "command", lint_tester.Classify("/soil:lint", roots))
	assert.String(t, "system", lint_tester.Classify("/api/goals", roots))
	assert.String(t, "system", lint_tester.Classify("/debug/pprof/", roots))
	assert.String(t, "pattern", lint_tester.Classify("//nolint", roots))
	assert.String(t, "system", lint_tester.Classify("/etc/hosts", roots))
	assert.String(t, "route", lint_tester.Classify("route:/api/goals", roots))
	assert.String(t, "path", lint_tester.Classify("path:/bin/true", roots))
	assert.String(t, "pattern", lint_tester.Classify("s/a/b/g", roots))
	assert.String(
		t,
		"import",
		lint_tester.Classify("\"example.org/module\"", roots),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify("pkg/web/RecoveryMiddleware", roots),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify("pkg/provision/salt.Client", roots),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify("pkg/check/memory.LocalLines()", roots),
	)
	assert.String(
		t,
		"symbol",
		lint_tester.Classify("go:pkg/provision/salt.Client", roots),
	)
	assert.String(
		t,
		"symbol",
		lint_tester.Classify(
			"go:../github/soil/pkg/provision/salt.Client",
			roots,
		),
	)
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify("go:pkg/<name>/Symbol", roots),
	)
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify("doc/ai/runbook/<name>.md", roots),
	)
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify("pkg/tool/*.go", roots),
	)
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify("$HOME/notes.md", roots),
	)
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify(
			"${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md",
			roots,
		),
	)
	assert.String(
		t,
		"sibling",
		lint_tester.Classify("../github/soil/doc/ai/spec/naming.md", roots),
	)
	assert.String(
		t,
		"absolute",
		lint_tester.Classify("/Users/example/notes.md", roots),
	)
	assert.String(
		t,
		"locator",
		lint_tester.Classify("https://code.claude.com/docs/en/skills", roots),
	)
}
