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

func TestVerdictLetsGitTagListingsThrough(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "git tag"))
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "git tag --sort=-v:refname | head -1"),
	)
	assert.String(t, "", guard.Verdict("darwin", "git tag -l 'v0.11.*'"))
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "git -C ../soil tag --contains HEAD"),
	)
	assert.String(t, "", guard.Verdict("linux", "git tag -n5 --merged main"))
}

func TestVerdictBlocksGitTagCreationAndDeletion(t *testing.T) {
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("darwin", "git tag v0.11.182"),
	)
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("darwin", "git tag -a v1.0.0 -m 'alfa'"),
	)
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("darwin", "git tag -d v0.11.181"),
	)
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("darwin", "git tag -lf v1.0.0"),
	)
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("darwin", "git tag \"$name\""),
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

func TestGitStashReadsPassUnasked(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "git stash list"))
	assert.String(t, "", guard.Question("git stash list"))
	assert.String(t, "", guard.Question("git -C ../soil stash list | head -1"))
	assert.String(t, "", guard.Question("git stash show -p stash@{1}"))
	assert.String(t, "", guard.Verdict("darwin", "git stash show \"$ref\""))
	assert.String(t, "", guard.Question("git stash show \"$ref\""))
}

func TestQuestionNamesTheGitStashVerb(t *testing.T) {
	assert.String(
		t,
		"git stash pop rewrites the shared index or work tree: everything it covers",
		guard.Question("git stash pop"),
	)
	assert.String(
		t,
		"git stash push rewrites the shared index or work tree: pkg/a.go",
		guard.Question("git stash push -- pkg/a.go"),
	)
}

func TestGitMoveDryRunPassesUnasked(t *testing.T) {
	assert.String(t, "", guard.Question("git mv -n pkg/a.go pkg/b.go"))
	assert.String(t, "", guard.Question("git mv --dry-run pkg/a.go pkg/b.go"))
	assert.String(
		t,
		"git mv rewrites the shared index or work tree: pkg/b.go",
		guard.Question("git mv -- -n pkg/b.go"),
	)
}

func TestVerdictReadsGitInsideAShellString(t *testing.T) {
	assert.String(
		t,
		"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("darwin", "bash -c 'git commit -m alfa'"),
	)
	assert.String(
		t,
		"git reset with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.Verdict("darwin", `sh -lc 'git reset -- $paths'`),
	)
	assert.String(
		t,
		"git push writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.Verdict("linux", "eval git push origin main"),
	)
	assert.String(
		t,
		"git rm rewrites the shared index or work tree: pkg/a.go",
		guard.Question("zsh -c 'cd soil && git rm pkg/a.go'"),
	)
}

func TestVerdictBlocksGitHiddenInAnUnreadableString(t *testing.T) {
	assert.String(
		t,
		"git inside a shell string or eval that cannot be read before it runs - run git directly, with every path written out",
		guard.Verdict("darwin", `bash -c "git reset -- $paths"`),
	)
	assert.String(
		t,
		"git inside a shell string or eval that cannot be read before it runs - run git directly, with every path written out",
		guard.Verdict("darwin", `eval "git $action"`),
	)
}

func TestVerdictLeavesShellStringsWithoutGit(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", `bash -c "echo $HOME"`))
	assert.String(t, "", guard.Verdict("darwin", "sh -c 'git status --short'"))
	assert.String(t, "", guard.Verdict("darwin", "ssh host 'git reset -- $x'"))
	assert.String(t, "", guard.Verdict("darwin", "bash"))
	assert.String(t, "", guard.Verdict("darwin", "sh script.sh -c"))
}

func TestAskDecisionShape(t *testing.T) {
	assert.String(
		t,
		`{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"ask","permissionDecisionReason":"alfa"}}`,
		string(notation.Marshal(guard.NewDecision("alfa", ""))),
	)
}
