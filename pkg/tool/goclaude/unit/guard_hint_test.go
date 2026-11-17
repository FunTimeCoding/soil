package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/response"
	"testing"
)

func TestTaskHintForDirectCodegen(t *testing.T) {
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
		guard.BlockReason(
			"darwin",
			"cd generated/client && oapi-codegen --config config.yaml ../server/openapi.yaml",
		),
	)
}

func TestScratchHintForTemporaryRedirect(t *testing.T) {
	assert.String(
		t,
		"scratch belongs under tmp/_session/<callsign>-<date>/ in the repository, where a later session can find it; the system temp is shared with every process and wiped on reboot",
		guard.Hint("go test ./... > /tmp/alfa.txt 2>&1"),
	)
	assert.StringContains(
		t,
		"tmp/_session",
		guard.Hint(
			"cd pkg && grep -rn bravo . >> /private/tmp/bravo-report.txt",
		),
	)
	assert.StringContains(
		t,
		"tmp/_session",
		guard.Hint(
			`for r in alfa bravo; do ls "$r" > "/tmp/charlie-$r.txt"; done`,
		),
	)
}

func TestHintEmptyForScratchAndDeviceRedirects(t *testing.T) {
	assert.String(
		t,
		"",
		guard.Hint("go test ./... > tmp/_session/alfa/test.txt"),
	)
	assert.String(t, "", guard.Hint("go vet ./... 2> /dev/null"))
	assert.String(t, "", guard.Hint("cat /tmp/alfa.txt"))
	assert.String(t, "", guard.Hint("wc -l < /tmp/alfa.txt"))
}

func TestHintEmptyWithoutCodegen(t *testing.T) {
	assert.String(t, "", guard.Hint("task generate"))
	assert.String(t, "", guard.Hint("grep -n oapi-codegen taskfile.yaml"))
	assert.String(t, "", guard.Hint("git status --short"))
}

func TestHintDecisionShape(t *testing.T) {
	assert.String(
		t,
		`{"hookSpecificOutput":{"hookEventName":"PreToolUse","additionalContext":"bravo"}}`,
		string(notation.Marshal(response.NewDecision("", "bravo"))),
	)
	assert.String(
		t,
		`{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"ask","permissionDecisionReason":"alfa","additionalContext":"bravo"}}`,
		string(notation.Marshal(response.NewDecision("alfa", "bravo"))),
	)
}
