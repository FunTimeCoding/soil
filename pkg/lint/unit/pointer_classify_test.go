package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"testing"
)

func TestClassifyRepository(t *testing.T) {
	assert.String(
		t,
		"repository",
		lint_tester.Classify("doc/ai/spec/naming.md"),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify(".claude/skills/sign-firefox/SKILL.md"),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify(".claude-plugin/plugin.json"),
	)
	assert.String(t, "repository", lint_tester.Classify("./doc/ai/spec"))
	assert.String(
		t,
		"repository",
		lint_tester.Classify("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/constant.md"),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify("pkg/web/RecoveryMiddleware"),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify("pkg/provision/salt.Client"),
	)
	assert.String(
		t,
		"repository",
		lint_tester.Classify("pkg/check/memory.LocalLines()"),
	)
}

func TestClassifyShort(t *testing.T) {
	assert.String(t, "short", lint_tester.Classify("tmp/gosec.json"))
	assert.String(t, "short", lint_tester.Classify(soil.SoilModule))
}

func TestClassifyCommand(t *testing.T) {
	assert.String(t, "command", lint_tester.Classify("/chart-sessions"))
	assert.String(t, "command", lint_tester.Classify("/soil:lint"))
}

func TestClassifySystem(t *testing.T) {
	assert.String(t, "system", lint_tester.Classify("/api/goals"))
	assert.String(t, "system", lint_tester.Classify("/debug/pprof/"))
	assert.String(t, "system", lint_tester.Classify("/etc/hosts"))
}

func TestClassifyPattern(t *testing.T) {
	assert.String(t, "pattern", lint_tester.Classify("//nolint"))
	assert.String(t, "pattern", lint_tester.Classify("s/a/b/g"))
}

func TestClassifyRoute(t *testing.T) {
	assert.String(t, "route", lint_tester.Classify("route:/api/goals"))
}

func TestClassifyPath(t *testing.T) {
	assert.String(t, "path", lint_tester.Classify("path:/bin/true"))
}

func TestClassifyMedia(t *testing.T) {
	assert.String(t, "media", lint_tester.Classify("media:text/html"))
	assert.String(t, "media", lint_tester.Classify("media:txt/html"))
}

func TestClassifyImport(t *testing.T) {
	assert.String(t, "import", lint_tester.Classify("\"example.org/module\""))
}

func TestClassifySymbol(t *testing.T) {
	assert.String(
		t,
		"symbol",
		lint_tester.Classify("go:pkg/provision/salt.Client"),
	)
	assert.String(
		t,
		"symbol",
		lint_tester.Classify("go:../github/soil/pkg/provision/salt.Client"),
	)
}

func TestClassifyPlaceholder(t *testing.T) {
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify("go:pkg/<name>/Symbol"),
	)
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify("doc/ai/runbook/<name>.md"),
	)
	assert.String(t, "placeholder", lint_tester.Classify("pkg/tool/*.go"))
	assert.String(t, "placeholder", lint_tester.Classify("$HOME/notes.md"))
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify("media:<type>/<subtype>"),
	)
	assert.String(
		t,
		"placeholder",
		lint_tester.Classify("${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md"),
	)
}

func TestClassifySibling(t *testing.T) {
	assert.String(
		t,
		"sibling",
		lint_tester.Classify("../github/soil/doc/ai/spec/naming.md"),
	)
}

func TestClassifyAbsolute(t *testing.T) {
	assert.String(
		t,
		"absolute",
		lint_tester.Classify("/Users/example/notes.md"),
	)
}

func TestClassifyLocator(t *testing.T) {
	assert.String(
		t,
		"locator",
		lint_tester.Classify("https://code.claude.com/docs/en/skills"),
	)
}
