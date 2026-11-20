package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/unit/grant_tester"
	"path/filepath"
	"testing"
	"time"
)

func TestGrantAllowsLabelledRepository(t *testing.T) {
	r := grant_tester.New(t)
	r.Label(r.Repository)
	resolved, e := filepath.EvalSymlinks(r.Repository)
	assert.FatalOnError(t, e)
	reason, notes := r.Check(
		r.Repository,
		"git add a.go && git commit -m change && git push origin main",
	)
	assert.String(t, "", reason)
	assert.Strings(
		t,
		[]string{
			fmt.Sprintf(
				"git add allowed by the git-write label for %s",
				resolved,
			),
			fmt.Sprintf(
				"git commit allowed by the git-write label for %s",
				resolved,
			),
			fmt.Sprintf(
				"git push allowed by the git-write label for %s",
				resolved,
			),
		},
		notes,
	)
}

func TestGrantKeyedBySession(t *testing.T) {
	r := grant_tester.New(t)
	r.Label(r.Repository)
	other := "other-session"
	r.Mock.Join(other, 1, time.Now())
	output, _, code := r.Guard(r.Session, r.Repository, "git commit -m change")
	assert.Integer(t, 0, code)
	assert.StringContains(
		t,
		"git commit allowed by the git-write label",
		output,
	)
	_, failure, code := r.Guard(other, r.Repository, "git commit -m change")
	assert.Integer(t, 2, code)
	assert.StringContains(
		t,
		"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
		failure,
	)
}

func TestGrantFollowsDirectoryOption(t *testing.T) {
	r := grant_tester.New(t)
	r.Label(r.Repository)
	reason, _ := r.Check(
		r.Other,
		fmt.Sprintf("git -C %s commit -m change", r.Repository),
	)
	assert.String(t, "", reason)
	reason, _ = r.Check(
		r.Repository,
		fmt.Sprintf("git -C %s commit -m change", r.Other),
	)
	assert.String(
		t,
		"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
		reason,
	)
}

func TestGrantRefusesOtherRepository(t *testing.T) {
	r := grant_tester.New(t)
	r.Label(r.Repository)
	reason, notes := r.Check(r.Other, "git commit -m change")
	assert.String(
		t,
		"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
		reason,
	)
	assert.Strings(t, nil, notes)
}

func TestGrantRefusesWithoutLabel(t *testing.T) {
	r := grant_tester.New(t)
	reason, _ := r.Check(r.Repository, "git push origin main")
	assert.String(
		t,
		"git push writes history - staging, committing and publishing belong to the user; leave it to them",
		reason,
	)
}

func TestGrantRefusesWritersOutsideSet(t *testing.T) {
	r := grant_tester.New(t)
	r.Label(r.Repository)

	for _, s := range []string{"tag", "merge", "rebase", "cherry-pick"} {
		reason, _ := r.Check(r.Repository, fmt.Sprintf("git %s main", s))
		assert.String(
			t,
			fmt.Sprintf(
				"git %s writes history - staging, committing and publishing belong to the user; leave it to them",
				s,
			),
			reason,
		)
	}
}

func TestGrantRefusesDangerousPush(t *testing.T) {
	r := grant_tester.New(t)
	r.Label(r.Repository)

	for _, command := range []string{
		"git push --force origin main",
		"git push -f origin main",
		"git push -uf origin main",
		"git push --force-with-lease=main origin main",
		"git push origin +main",
		"git push origin :main",
		"git push --tags origin",
		"git push --follow-tags origin main",
		"git push --delete origin main",
		"git push --mirror origin",
		"git push --all origin",
		"git push --repo=elsewhere",
		"git push https://example.com/other.git main",
		"git push ../other main",
	} {
		reason, _ := r.Check(r.Repository, command)
		assert.String(
			t,
			"git push writes history - staging, committing and publishing belong to the user; leave it to them",
			reason,
		)
	}
}

func TestGrantRefusesUnclearTarget(t *testing.T) {
	r := grant_tester.New(t)
	r.Label(r.Repository)

	for _, command := range []string{
		fmt.Sprintf("cd %s && git commit -m change", r.Repository),
		"GIT_DIR=/elsewhere git commit -m change",
		"export GIT_WORK_TREE=/elsewhere; git commit -m change",
		"git -c core.worktree=/elsewhere commit -m change",
		"git --git-dir=/elsewhere commit -m change",
		"git --git-dir /elsewhere commit -m change",
		`git -C "$target" commit -m change`,
		"git -C ~/elsewhere commit -m change",
		`git commit -m change "$file"`,
		"bash -c 'git commit -m change'",
	} {
		reason, _ := r.Check(r.Repository, command)
		assert.String(
			t,
			"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
			reason,
		)
	}
}

func TestGrantRefusesWhenDaemonUnreachable(t *testing.T) {
	r := grant_tester.New(t)
	r.Label(r.Repository)
	r.Mock.Fail(errors.New("connection refused"))
	reason, _ := r.Check(r.Repository, "git commit -m change")
	assert.String(
		t,
		"git commit writes history, and the git-write label could not be read from goclauded (connection refused) - leave it to the user",
		reason,
	)
}

func TestGitDirectoryOptionValueSkipped(t *testing.T) {
	assert.String(
		t,
		"git commit writes history - staging, committing and publishing belong to the user; leave it to them",
		guard.BlockReason("darwin", "git --git-dir /elsewhere commit -m change"),
	)
}
