//go:build local

package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/repository_tester"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohook/integration/gohook_tester"
	"path/filepath"
	"testing"
)

func TestPreCommitThroughGitRefusesFixer(t *testing.T) {
	clone := gohook_tester.Installed(
		t,
		`pre-commit:
  - paths: ['**/*.go']
    run: printf 'package a // fixed\n' > a.go
`,
	)
	repository_tester.NewAt(clone).Write("a.go", "package a\n")
	repository_tester.Command(clone, "add", "a.go")
	r := gohook_tester.Attempt(
		clone,
		nil,
		"commit",
		"-q",
		gitConstant.MessageArgument,
		"first",
	)
	assert.Integer(t, 1, r.Exit)
	assert.StringContains(
		t,
		"modified files that are not staged",
		r.ErrorString,
	)
	assert.StringContains(t, "a.go", r.ErrorString)
	assert.String(
		t,
		"package a // fixed\n",
		system.ReadFileUnsafe(filepath.Join(clone, "a.go")),
	)
	repository_tester.Command(clone, "add", "a.go")
	assert.Integer(
		t,
		0,
		gohook_tester.Attempt(
			clone,
			nil,
			"commit",
			"-q",
			gitConstant.MessageArgument,
			"first",
		).Exit,
	)
}

func TestCommitMessageThroughGitReceivesArgument(t *testing.T) {
	clone := gohook_tester.Installed(
		t,
		"commit-msg:\n  - run: grep -q '^[a-z]' {1}\n",
	)
	repository_tester.NewAt(clone).Write("b.go", "package b\n")
	repository_tester.Command(clone, "add", "b.go")
	r := gohook_tester.Attempt(
		clone,
		nil,
		"commit",
		"-q",
		gitConstant.MessageArgument,
		"Capitalised",
	)
	assert.Integer(t, 1, r.Exit)
	assert.StringContains(t, "job 1 failed", r.ErrorString)
	assert.Integer(
		t,
		0,
		gohook_tester.Attempt(
			clone,
			nil,
			"commit",
			"-q",
			gitConstant.MessageArgument,
			"lowercase",
		).Exit,
	)
}

func TestSkipVariableThroughGit(t *testing.T) {
	clone := gohook_tester.Installed(t, "pre-commit:\n  - run: exit 1\n")
	repository_tester.NewAt(clone).Write("c.go", "package c\n")
	repository_tester.Command(clone, "add", "c.go")
	assert.Integer(
		t,
		1,
		gohook_tester.Attempt(
			clone,
			nil,
			"commit",
			"-q",
			gitConstant.MessageArgument,
			"blocked",
		).Exit,
	)
	assert.Integer(
		t,
		0,
		gohook_tester.Attempt(
			clone,
			map[string]string{constant.SkipEnvironment: constant.SkipValue},
			"commit",
			"-q",
			gitConstant.MessageArgument,
			"skipped",
		).Exit,
	)
}

func TestPrePushThroughGitFiltersByPath(t *testing.T) {
	clone := gohook_tester.Installed(
		t,
		`pre-push:
  - paths: [doc/]
    run: touch ran-doc
  - paths: [pkg/]
    run: touch ran-pkg
`,
	)
	repository_tester.NewAt(clone).Write("doc/x.md", "x\n")
	repository_tester.NewAt(clone).Commit("doc only")
	assert.Integer(t, 0, gohook_tester.Attempt(clone, nil, "push", "-q").Exit)
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-doc")))
	assert.False(t, system.FileExists(filepath.Join(clone, "ran-pkg")))
}
