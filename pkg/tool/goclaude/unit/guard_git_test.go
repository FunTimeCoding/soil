package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestVerdictBlocksAGitWriterWithAnExpandedPath(t *testing.T) {
	assert.String(
		t,
		"git reset with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.Verdict(
			"darwin",
			"paths=$(git status --short) && git reset -q -- $paths",
		),
	)
	assert.String(
		t,
		"git checkout with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.Verdict("linux", "git checkout -- *.go"),
	)
	assert.String(
		t,
		"git rm with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.Verdict("darwin", "find . -name '*.orig' | xargs git rm"),
	)
	assert.String(
		t,
		"git restore with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.Verdict("darwin", "git restore --staged --"),
	)
	assert.String(
		t,
		"git checkout with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.Verdict("darwin", "xargs git checkout -- < list.txt"),
	)
}

func TestVerdictBlocksGitHistoryWriters(t *testing.T) {
	assert.String(
		t,
		"git add writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("darwin", "git add -N pkg/a.go"),
	)
	assert.String(
		t,
		"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict(
			"darwin",
			"git -C /tmp/repo -c user.name=x commit -m 'alfa'",
		),
	)
	assert.String(
		t,
		"git push writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("linux", "cd soil && git push origin main"),
	)
}

func TestVerdictLetsGitReadersThrough(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "git status --short -- pkg"))
	assert.String(t, "", guard.Verdict("darwin", "git diff --cached --stat"))
	assert.String(t, "", guard.Verdict("darwin", "git -C ../soil log -3"))
	assert.String(t, "", guard.Question("git status --short -- pkg"))
}

func TestQuestionAsksForALiteralGitTreeWriter(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "git mv pkg/a.go pkg/b.go"))
	assert.String(
		t,
		"git mv rewrites the shared index or work tree: pkg/a.go pkg/b.go",
		guard.Question("git mv pkg/a.go pkg/b.go"),
	)
	assert.String(
		t,
		"git reset rewrites the shared index or work tree: pkg/a.go",
		guard.Question("git reset -q -- pkg/a.go"),
	)
	assert.String(
		t,
		"git stash rewrites the shared index or work tree: everything it covers",
		guard.Question("git stash"),
	)
}

func TestAskDecisionShape(t *testing.T) {
	assert.String(
		t,
		`{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"ask","permissionDecisionReason":"alfa"}}`,
		string(notation.Marshal(guard.NewAskDecision("alfa"))),
	)
}
