package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/response"
	"testing"
)

func TestGitWriterWithExpandedPathBlocked(t *testing.T) {
	assert.String(
		t,
		"git reset with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason(
			"darwin",
			"paths=$(git status --short) && git reset -q -- $paths",
		),
	)
	assert.String(
		t,
		"git checkout with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("linux", "git checkout -- *.go"),
	)
	assert.String(
		t,
		"git checkout with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("linux", "git checkout -- {a,b}.go"),
	)
	assert.String(
		t,
		"git restore with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("linux", `git restore a\ b.go`),
	)
	assert.String(
		t,
		"git rm with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("linux", `git rm $'a.go'`),
	)
	assert.String(
		t,
		"git checkout with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("linux", "xargs -I {} git checkout -- {}"),
	)
	assert.String(
		t,
		"git rm with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("darwin", "find . -name '*.orig' | xargs git rm"),
	)
	assert.String(
		t,
		"git restore with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("darwin", "git restore --staged --"),
	)
	assert.String(
		t,
		"git checkout with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("darwin", "xargs git checkout -- < list.txt"),
	)
}

func TestGitHistoryWritersBlocked(t *testing.T) {
	assert.String(
		t,
		"git add writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("darwin", "git add -N pkg/a.go"),
	)
	assert.String(
		t,
		"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason(
			"darwin",
			"git -C /tmp/repo -c user.name=x commit -m 'alfa'",
		),
	)
	assert.String(
		t,
		"git push writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("linux", "cd soil && git push origin main"),
	)
}

func TestGitTagListingsAllowed(t *testing.T) {
	assert.String(t, "", guard.BlockReason("darwin", "git tag"))
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", "git tag --sort=-v:refname | head -1"),
	)
	assert.String(t, "", guard.BlockReason("darwin", "git tag -l 'v0.11.*'"))
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", "git -C ../soil tag --contains HEAD"),
	)
	assert.String(
		t,
		"",
		guard.BlockReason("linux", "git tag -n5 --merged main"),
	)
}

func TestGitTagCreationAndDeletionBlocked(t *testing.T) {
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("darwin", "git tag v0.11.182"),
	)
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("darwin", "git tag -a v1.0.0 -m 'alfa'"),
	)
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("darwin", "git tag -d v0.11.181"),
	)
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("darwin", "git tag -lf v1.0.0"),
	)
	assert.String(
		t,
		"git tag writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("darwin", "git tag \"$name\""),
	)
}

func TestGitReadersAllowed(t *testing.T) {
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", "git status --short -- pkg"),
	)
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", "git diff --cached --stat"),
	)
	assert.String(t, "", guard.BlockReason("darwin", "git -C ../soil log -3"))
	assert.String(t, "", guard.AskReason("git status --short -- pkg"))
}

func TestWrappedGitTreeWriterNeedsApproval(t *testing.T) {
	assert.String(
		t,
		"git checkout rewrites the shared index or work tree: a.go",
		guard.AskReason("timeout 5 git checkout -- a.go"),
	)
}

func TestLiteralGitTreeWriterNeedsApproval(t *testing.T) {
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", "git mv pkg/a.go pkg/b.go"),
	)
	assert.String(
		t,
		"git mv rewrites the shared index or work tree: pkg/a.go pkg/b.go",
		guard.AskReason("git mv pkg/a.go pkg/b.go"),
	)
	assert.String(
		t,
		"git reset rewrites the shared index or work tree: pkg/a.go",
		guard.AskReason("git reset -q -- pkg/a.go"),
	)
	assert.String(
		t,
		"git stash rewrites the shared index or work tree: everything it covers",
		guard.AskReason("git stash"),
	)
}

func TestGitStashReadsPassUnasked(t *testing.T) {
	assert.String(t, "", guard.BlockReason("darwin", "git stash list"))
	assert.String(t, "", guard.AskReason("git stash list"))
	assert.String(t, "", guard.AskReason("git -C ../soil stash list | head -1"))
	assert.String(t, "", guard.AskReason("git stash show -p stash@{1}"))
	assert.String(t, "", guard.BlockReason("darwin", "git stash show \"$ref\""))
	assert.String(t, "", guard.AskReason("git stash show \"$ref\""))
}

func TestGitStashAskReasonWithVerb(t *testing.T) {
	assert.String(
		t,
		"git stash pop rewrites the shared index or work tree: everything it covers",
		guard.AskReason("git stash pop"),
	)
	assert.String(
		t,
		"git stash push rewrites the shared index or work tree: pkg/a.go",
		guard.AskReason("git stash push -- pkg/a.go"),
	)
}

func TestGitMoveDryRunPassesUnasked(t *testing.T) {
	assert.String(t, "", guard.AskReason("git mv -n pkg/a.go pkg/b.go"))
	assert.String(t, "", guard.AskReason("git mv --dry-run pkg/a.go pkg/b.go"))
	assert.String(
		t,
		"git mv rewrites the shared index or work tree: pkg/b.go",
		guard.AskReason("git mv -- -n pkg/b.go"),
	)
}

func TestGitInsideShellStringBlocked(t *testing.T) {
	assert.String(
		t,
		"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("darwin", "bash -c 'git commit -m alfa'"),
	)
	assert.String(
		t,
		"git reset with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally",
		guard.BlockReason("darwin", `sh -lc 'git reset -- $paths'`),
	)
	assert.String(
		t,
		"git push writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("linux", "eval git push origin main"),
	)
	assert.String(
		t,
		"git rm rewrites the shared index or work tree: pkg/a.go",
		guard.AskReason("zsh -c 'cd soil && git rm pkg/a.go'"),
	)
}

func TestGitHiddenInUnreadableStringBlocked(t *testing.T) {
	assert.String(
		t,
		"git inside a shell string or eval that cannot be read before it runs - run git directly, with every path written out",
		guard.BlockReason("darwin", `bash -c "git reset -- $paths"`),
	)
	assert.String(
		t,
		"git inside a shell string or eval that cannot be read before it runs - run git directly, with every path written out",
		guard.BlockReason("darwin", `eval "git $action"`),
	)
	assert.String(
		t,
		"git inside a shell string or eval that cannot be read before it runs - run git directly, with every path written out",
		guard.BlockReason("darwin", `eval git "$action"`),
	)
}

func TestShellStringsWithoutGitAllowed(t *testing.T) {
	assert.String(t, "", guard.BlockReason("darwin", `bash -c "echo $HOME"`))
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", "sh -c 'git status --short'"),
	)
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", "ssh host 'git reset -- $x'"),
	)
	assert.String(t, "", guard.BlockReason("darwin", "bash"))
	assert.String(t, "", guard.BlockReason("darwin", "sh script.sh -c"))
}

func TestAskDecisionShape(t *testing.T) {
	assert.String(
		t,
		`{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"ask","permissionDecisionReason":"alfa"}}`,
		string(notation.Marshal(response.NewDecision("alfa", ""))),
	)
}
