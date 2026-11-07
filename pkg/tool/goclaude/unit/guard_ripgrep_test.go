package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestQuestionLetsRipgrepSearchesThrough(t *testing.T) {
	assert.String(t, "", guard.Question("rg -n 'func New' pkg"))
	assert.String(t, "", guard.Question(`rg -l "\b$b\(" --type go .`))
	assert.String(t, "", guard.Question(`rg -q foo -- "$f" --pre`))
	assert.String(t, "", guard.Question("rg --pre-glob '*.gz' --no-pre foo"))
	assert.String(t, "", guard.Question("git ls-files | rg -c foo"))
	assert.String(t, "", guard.Question("timeout 5 rg foo pkg/a.go"))
	assert.String(t, "", guard.Question("command -v rg"))
	assert.String(t, "", guard.Question("xargs grep -l rg < list"))
	assert.String(t, "", guard.Question("grep -l foo pkg && ("))
}

func TestQuestionAsksForARipgrepProgram(t *testing.T) {
	assert.String(
		t,
		"rg --pre runs a program",
		guard.Question("rg --pre gunzip foo ."),
	)
	assert.String(
		t,
		"rg --pre=cat runs a program",
		guard.Question("n=$(rg -n --'pre'=cat foo)"),
	)
	assert.String(
		t,
		"rg --hostname-bin=/bin/hostname runs a program",
		guard.Question("rg --hostname-bin=/bin/hostname foo"),
	)
	assert.String(
		t,
		"rg --pre runs a program",
		guard.Question("nice -n 5 timeout -s KILL 5 rg --pre cat foo"),
	)
}

func TestQuestionAsksForUnreadableRipgrepArguments(t *testing.T) {
	reason := "rg with an argument from an expansion, a glob or xargs input, or with a config file from RIPGREP_CONFIG_PATH, could run a program through --pre"

	for _, command := range []string{
		"rg -q foo $f",
		`rg foo "$f"`,
		"rg foo *",
		"rg foo {--pre=cat,x}",
		"rg --pr{e,}=cat foo",
		`rg $'\x2d-pre' cat foo`,
		`rg --$'\x70re' cat foo`,
		`rg --\pre cat foo`,
		"find . -name '*.go' | xargs rg foo",
		"xargs -n 1 rg foo < list",
		"RIPGREP_CONFIG_PATH=tmp/x rg foo",
		"export RIPGREP_CONFIG_PATH=tmp/x; rg foo",
		"rg foo pkg && (",
		`rg foo pkg; bash -c 'if'`,
	} {
		assert.String(t, reason, guard.Question(command))
	}
}

func TestQuestionJoinsGitAndRipgrepReasons(t *testing.T) {
	assert.String(
		t,
		"git mv rewrites the shared index or work tree: pkg/a.go pkg/b.go; rg --pre runs a program",
		guard.Question(
			"git mv pkg/a.go pkg/b.go && rg --pre cat x && rg --pre cat y",
		),
	)
}
