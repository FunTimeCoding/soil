package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/changed"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohook/option"
	"path/filepath"
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
	command(root, "init", "--initial-branch=main", library.CurrentDirectory)
	assert.String(t, "", changed.Upstream(root))
	assert.True(t, changed.NewPush(root).All)
}

func TestInstallRefusesRedirectedHooksPath(t *testing.T) {
	clone := repository(t)
	write(clone, constant.RootFile, "pre-push:\n  - run: a\n")
	command(clone, "config", "core.hooksPath", "/dev/null")
	assert.String(t, "/dev/null", git.HooksPath(clone))
	defer func() { assert.NotNil(t, recover()) }()
	gohook.Install(configuration.Load(clone))
}

func TestInstallAndUninstall(t *testing.T) {
	clone := repository(t)
	write(
		clone,
		constant.ToolFile,
		"pre-push:\n  - run: a\npre-commit:\n  - run: b\n",
	)
	c := configuration.Load(clone)
	foreign := filepath.Join(c.HooksDirectory(), "post-merge")
	write(c.HooksDirectory(), "post-merge", "#!/bin/sh\necho theirs\n")
	gohook.Install(c)
	stub := filepath.Join(c.HooksDirectory(), "pre-push")
	assert.True(t, system.FileExists(stub))
	assert.True(
		t,
		system.FileExists(filepath.Join(c.HooksDirectory(), "pre-commit")),
	)
	assert.StringContains(
		t,
		"exec gohook run pre-push",
		system.ReadFileUnsafe(stub),
	)
	assert.True(t, system.IsExecutable(stub))
	gohook.Uninstall(c)
	assert.False(t, system.FileExists(stub))
	assert.True(t, system.FileExists(foreign))
}

func TestRunPreCommitUsesStagedFiles(t *testing.T) {
	clone := repository(t)
	write(
		clone,
		constant.RootFile,
		`pre-commit:
  - paths: ['**/*.go']
    run: touch ran-go
  - paths: [doc/]
    run: touch ran-doc
`,
	)
	write(clone, "a.go", "package a\n")
	write(clone, "doc/x.md", "x\n")
	command(clone, "add", "a.go")
	o := option.New()
	o.Hook = "pre-commit"
	gohook.Run(configuration.Load(clone), o)
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-go")))
	assert.False(t, system.FileExists(filepath.Join(clone, "ran-doc")))
}

func TestRunCommitMessagePassesArgument(t *testing.T) {
	clone := repository(t)
	write(
		clone,
		constant.RootFile,
		"commit-msg:\n  - run: cp {1} seen-message\n",
	)
	write(clone, "message.txt", "subject\n")
	o := option.New()
	o.Hook = "commit-msg"
	o.Arguments = []string{"message.txt"}
	gohook.Run(configuration.Load(clone), o)
	assert.String(
		t,
		"subject\n",
		system.ReadFileUnsafe(filepath.Join(clone, "seen-message")),
	)
}

func TestUnstagedDetectsFixerOutput(t *testing.T) {
	clone := repository(t)
	write(clone, "a.go", "package a\n")
	commit(clone, "add")
	assert.Count(t, 0, changed.Unstaged(clone))
	write(clone, "a.go", "package a // fixed\n")
	assert.Strings(t, []string{"a.go"}, changed.Unstaged(clone))
}

func TestRunPrePushFiltersJobs(t *testing.T) {
	clone := repository(t)
	write(
		clone,
		constant.ToolFile,
		`pre-push:
  - paths: ['**/*.go']
    run: touch ran-go
  - paths: [strata/manifest/]
    run: touch ran-manifest
  - run: touch ran-always
`,
	)
	write(clone, "pkg/a.go", "package a\n")
	commit(clone, "go change")
	o := option.New()
	o.Hook = "pre-push"
	gohook.Run(configuration.Load(clone), o)
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-go")))
	assert.False(t, system.FileExists(filepath.Join(clone, "ran-manifest")))
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-always")))
}

func TestRunExplicitRange(t *testing.T) {
	clone := repository(t)
	write(
		clone,
		constant.RootFile,
		"pre-push:\n  - paths: [doc/]\n    run: touch ran-doc\n",
	)
	write(clone, "doc/x.md", "x\n")
	commit(clone, "doc")
	o := option.New()
	o.Hook = "pre-push"
	o.Base = "HEAD~1"
	o.Head = "HEAD"
	gohook.Run(configuration.Load(clone), o)
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-doc")))
}

func TestRunHookWithoutJobs(t *testing.T) {
	clone := repository(t)
	write(clone, constant.RootFile, "pre-push:\n  - run: touch never\n")
	o := option.New()
	o.Hook = "post-merge"
	gohook.Run(configuration.Load(clone), o)
	assert.False(t, system.FileExists(filepath.Join(clone, "never")))
}
