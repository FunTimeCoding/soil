package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestHintPointsDirectCodegenAtTheTask(t *testing.T) {
	assert.String(
		t,
		"`task generate` in the repository root regenerates every service's client and server from its openapi.yaml with this same oapi-codegen invocation; it is how this repository runs code generation, and it runs quietly",
		guard.Hint(
			"cd pkg/tool/goclauded/generated/client && oapi-codegen --config config.yaml ../server/openapi.yaml",
		),
	)
	assert.StringContains(
		t,
		"task generate",
		guard.Hint(
			"sh -c 'cd generated/server && oapi-codegen --config config.yaml openapi.yaml'",
		),
	)
	assert.String(
		t,
		"",
		guard.Verdict(
			"darwin",
			"cd generated/client && oapi-codegen --config config.yaml ../server/openapi.yaml",
		),
	)
}

func TestHintStaysQuietWithoutCodegen(t *testing.T) {
	assert.String(t, "", guard.Hint("task generate"))
	assert.String(t, "", guard.Hint("grep -n oapi-codegen taskfile.yaml"))
	assert.String(t, "", guard.Hint("git status --short"))
}

func TestHintDecisionShape(t *testing.T) {
	assert.String(
		t,
		`{"hookSpecificOutput":{"hookEventName":"PreToolUse","additionalContext":"bravo"}}`,
		string(notation.Marshal(guard.NewDecision("", "bravo"))),
	)
	assert.String(
		t,
		`{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"ask","permissionDecisionReason":"alfa","additionalContext":"bravo"}}`,
		string(notation.Marshal(guard.NewDecision("alfa", "bravo"))),
	)
}
