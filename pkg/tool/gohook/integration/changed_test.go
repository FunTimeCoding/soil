package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/changed"
	"testing"
)

func TestPushRange(t *testing.T) {
	clone := repository(t)
	write(clone, "pkg/a.go", "package a\n")
	write(clone, "doc/x.md", "x\n")
	commit(clone, "work")
	r := changed.NewPush(clone)
	assert.False(t, r.All)
	assert.String(t, "origin/main..HEAD", r.String())
	assert.Strings(t, []string{"doc/x.md", "pkg/a.go"}, r.Files(clone))
}

func TestPushRangeNothingToPush(t *testing.T) {
	clone := repository(t)
	assert.Count(t, 0, changed.NewPush(clone).Files(clone))
}

func TestStagedRange(t *testing.T) {
	clone := repository(t)
	write(clone, "staged.go", "package s\n")
	write(clone, "unstaged.go", "package u\n")
	command(clone, "add", "staged.go")
	r := changed.NewStaged()
	assert.String(t, "staged files", r.String())
	assert.Strings(t, []string{"staged.go"}, r.Files(clone))
}

func TestHookRanges(t *testing.T) {
	clone := repository(t)
	assert.False(t, changed.NewHook(clone, "pre-push").All)
	assert.True(t, changed.NewHook(clone, "pre-commit").Staged)
	assert.True(t, changed.NewHook(clone, "post-merge").All)
	assert.String(t, "all files", changed.NewAll().String())
}

func TestExplicitRange(t *testing.T) {
	clone := repository(t)
	write(clone, "one.go", "package one\n")
	commit(clone, "one")
	write(clone, "two.go", "package two\n")
	commit(clone, "two")
	assert.Strings(
		t,
		[]string{"two.go"},
		changed.New("HEAD~1", "HEAD").Files(clone),
	)
}

func TestUpstreamWithoutTracking(t *testing.T) {
	root := t.TempDir()
	command(root, "init", "--initial-branch=main", constant.CurrentDirectory)
	assert.String(t, "", changed.Upstream(root))
	assert.True(t, changed.NewPush(root).All)
}
