package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"testing"
)

func TestClassifyRepository(t *testing.T) {
	assert.String(t, "repository", classify("doc/ai/spec/naming.md"))
	assert.String(
		t,
		"repository",
		classify(".claude/skills/sign-firefox/SKILL.md"),
	)
	assert.String(t, "repository", classify(".claude-plugin/plugin.json"))
	assert.String(t, "repository", classify("./doc/ai/spec"))
	assert.String(
		t,
		"repository",
		classify("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/constant.md"),
	)
	assert.String(t, "repository", classify("pkg/web/RecoveryMiddleware"))
	assert.String(t, "repository", classify("pkg/provision/salt.Client"))
	assert.String(t, "repository", classify("pkg/check/memory.LocalLines()"))
}

func TestClassifyShort(t *testing.T) {
	assert.String(t, "short", classify("tmp/gosec.json"))
	assert.String(t, "short", classify(constant.SoilModule))
}

func TestClassifyCommand(t *testing.T) {
	assert.String(t, "command", classify("/chart-sessions"))
	assert.String(t, "command", classify("/soil:lint"))
}

func TestClassifySystem(t *testing.T) {
	assert.String(t, "system", classify("/api/goals"))
	assert.String(t, "system", classify("/debug/pprof/"))
	assert.String(t, "system", classify("/etc/hosts"))
}

func TestClassifyPattern(t *testing.T) {
	assert.String(t, "pattern", classify("//nolint"))
	assert.String(t, "pattern", classify("s/a/b/g"))
}

func TestClassifyRoute(t *testing.T) {
	assert.String(t, "route", classify("route:/api/goals"))
}

func TestClassifyPath(t *testing.T) {
	assert.String(t, "path", classify("path:/bin/true"))
}

func TestClassifyImport(t *testing.T) {
	assert.String(t, "import", classify("\"example.org/module\""))
}

func TestClassifySymbol(t *testing.T) {
	assert.String(t, "symbol", classify("go:pkg/provision/salt.Client"))
	assert.String(
		t,
		"symbol",
		classify("go:../github/soil/pkg/provision/salt.Client"),
	)
}

func TestClassifyPlaceholder(t *testing.T) {
	assert.String(t, "placeholder", classify("go:pkg/<name>/Symbol"))
	assert.String(t, "placeholder", classify("doc/ai/runbook/<name>.md"))
	assert.String(t, "placeholder", classify("pkg/tool/*.go"))
	assert.String(t, "placeholder", classify("$HOME/notes.md"))
	assert.String(
		t,
		"placeholder",
		classify("${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md"),
	)
}

func TestClassifySibling(t *testing.T) {
	assert.String(
		t,
		"sibling",
		classify("../github/soil/doc/ai/spec/naming.md"),
	)
}

func TestClassifyAbsolute(t *testing.T) {
	assert.String(t, "absolute", classify("/Users/example/notes.md"))
}

func TestClassifyLocator(t *testing.T) {
	assert.String(
		t,
		"locator",
		classify("https://code.claude.com/docs/en/skills"),
	)
}
