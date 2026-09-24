package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"path/filepath"
	"testing"
)

func TestPreCommitThroughGitRefusesFixer(t *testing.T) {
	clone := installed(
		t,
		`pre-commit:
  - paths: ['**/*.go']
    run: printf 'package a // fixed\n' > a.go
`,
	)
	write(clone, "a.go", "package a\n")
	command(clone, "add", "a.go")
	r := attempt(clone, nil, "commit", "-q", "-m", "first")
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
	command(clone, "add", "a.go")
	assert.Integer(
		t,
		0,
		attempt(clone, nil, "commit", "-q", "-m", "first").Exit,
	)
}

func TestCommitMessageThroughGitReceivesArgument(t *testing.T) {
	clone := installed(t, "commit-msg:\n  - run: grep -q '^[a-z]' {1}\n")
	write(clone, "b.go", "package b\n")
	command(clone, "add", "b.go")
	r := attempt(clone, nil, "commit", "-q", "-m", "Capitalised")
	assert.Integer(t, 1, r.Exit)
	assert.StringContains(t, "job 1 failed", r.ErrorString)
	assert.Integer(
		t,
		0,
		attempt(clone, nil, "commit", "-q", "-m", "lowercase").Exit,
	)
}

func TestSkipVariableThroughGit(t *testing.T) {
	clone := installed(t, "pre-commit:\n  - run: exit 1\n")
	write(clone, "c.go", "package c\n")
	command(clone, "add", "c.go")
	assert.Integer(
		t,
		1,
		attempt(clone, nil, "commit", "-q", "-m", "blocked").Exit,
	)
	assert.Integer(
		t,
		0,
		attempt(
			clone,
			map[string]string{constant.SkipEnvironment: constant.SkipValue},
			"commit",
			"-q",
			"-m",
			"skipped",
		).Exit,
	)
}

func TestPrePushThroughGitFiltersByPath(t *testing.T) {
	clone := installed(
		t,
		`pre-push:
  - paths: [doc/]
    run: touch ran-doc
  - paths: [pkg/]
    run: touch ran-pkg
`,
	)
	write(clone, "doc/x.md", "x\n")
	commit(clone, "doc only")
	assert.Integer(t, 0, attempt(clone, nil, "push", "-q").Exit)
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-doc")))
	assert.False(t, system.FileExists(filepath.Join(clone, "ran-pkg")))
}
