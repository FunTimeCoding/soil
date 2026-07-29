package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestClassify(t *testing.T) {
	roots := []string{".claude", ".claude-plugin", "doc", "pkg", "skills"}
	assert.String(
		t,
		constant.PointerRepository,
		pointer.Classify(
			"doc/ai/spec/naming.md",
			roots,
		),
	)
	assert.String(
		t,
		constant.PointerRepository,
		pointer.Classify(
			".claude/skills/sign-firefox/SKILL.md",
			roots,
		),
	)
	assert.String(
		t,
		constant.PointerRepository,
		pointer.Classify(".claude-plugin/plugin.json", roots),
	)
	assert.String(
		t,
		constant.PointerRepository,
		pointer.Classify("./doc/ai/spec", roots),
	)
	assert.String(
		t,
		constant.PointerRepository,
		pointer.Classify(
			"${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/constant.md",
			roots,
		),
	)
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify("tmp/gosec.json", roots),
	)
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify(library.SoilModule, roots),
	)
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify("/chart-sessions", roots),
	)
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify("/api/goals", roots),
	)
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify("/debug/pprof/", roots),
	)
	assert.String(t, constant.PointerUnknown, pointer.Classify("//nolint", roots))
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify("/etc/hosts", roots),
	)
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify(
			"pkg/web/RecoveryMiddleware",
			roots,
		),
	)
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify(
			"pkg/provision/salt.Client",
			roots,
		),
	)
	assert.String(
		t,
		constant.PointerUnknown,
		pointer.Classify("pkg/check/memory.LocalLines()", roots),
	)
	assert.String(
		t,
		constant.PointerPlaceholder,
		pointer.Classify(
			"doc/ai/runbook/<name>.md",
			roots,
		),
	)
	assert.String(
		t,
		constant.PointerPlaceholder,
		pointer.Classify("pkg/tool/*.go", roots),
	)
	assert.String(
		t,
		constant.PointerPlaceholder,
		pointer.Classify("$HOME/notes.md", roots),
	)
	assert.String(
		t,
		constant.PointerPlaceholder,
		pointer.Classify(
			"${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md",
			roots,
		),
	)
	assert.String(
		t,
		constant.PointerSibling,
		pointer.Classify("../github/soil/doc/ai/spec/naming.md", roots),
	)
	assert.String(
		t,
		constant.PointerAbsolute,
		pointer.Classify(
			"/Users/example/notes.md",
			roots,
		),
	)
	assert.String(
		t,
		constant.PointerLocator,
		pointer.Classify("https://code.claude.com/docs/en/skills", roots),
	)
}
