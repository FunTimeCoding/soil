package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/changed"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohook/option"
	"path/filepath"
	"testing"
)

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
