package integration

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/command"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestRunStart(t *testing.T) {
	r1 := run.New()
	assert.True(t, r1.Panic)
	r1.Panic = false
	r1.Start("echo", "test")
	assert.True(t, r1.Error == nil)
	assert.String(t, "test\n", r1.OutputString)
	assert.String(t, "", r1.ErrorString)
	assert.Integer(t, 0, r1.Exit)
	r2 := run.New()
	r2.Panic = false
	r2.Start("nonexistent")
	assert.True(t, r2.Error != nil)
	assert.String(t, "", r2.OutputString)
	assert.String(t, "", r2.ErrorString)

	switch runtime.GOOS {
	case constant.Windows:
		assert.String(
			t,
			`nonexistent: exec: "nonexistent": executable file not found in %PATH%`,
			r2.Error.Error(),
		)
	default:
		assert.String(
			t,
			`nonexistent: exec: "nonexistent": executable file not found in $PATH`,
			r2.Error.Error(),
		)
	}
}

func TestStartFailureCarriesCommandError(t *testing.T) {
	windowsSkip(t)
	c := run.New().NoPanic()
	c.Start(constant.Shell, constant.ShellCommand, "echo oops >&2; exit 3")
	assert.NotNil(t, c.Error)
	assert.Integer(t, 3, c.Exit)
	assert.True(t, command.Is(c.Error))
	failure := c.Error.(*command.CommandError)
	assert.String(
		t,
		"/bin/sh -c echo oops >&2; exit 3: exit status 3",
		failure.Error(),
	)
	assert.StringContains(t, "oops", failure.Stderr)
}

func TestRunPanicMode(t *testing.T) {
	windowsSkip(t)
	var recovered any
	func() {
		defer func() { recovered = recover() }()
		run.New().Start(constant.Shell, constant.ShellCommand, "exit 7")
	}()
	assert.NotNil(t, recovered)
}

func TestRunDirectory(t *testing.T) {
	windowsSkip(t)
	directory, e := filepath.EvalSymlinks(t.TempDir())
	errors.PanicOnError(e)
	r := run.New()
	r.Directory = directory
	output := r.Start(constant.Shell, constant.ShellCommand, "pwd")
	assert.String(t, directory, strings.TrimSpace(output))
}

func TestRunEnvironment(t *testing.T) {
	windowsSkip(t)
	r := run.New()
	r.Environment("INTEGRATION_VALUE", "expected")
	output := r.Start(
		constant.Shell,
		constant.ShellCommand,
		"echo $INTEGRATION_VALUE",
	)
	assert.String(t, "expected\n", output)
}

func TestRunSetEnvironment(t *testing.T) {
	windowsSkip(t)
	r := run.New()
	r.SetEnvironment([]string{"INTEGRATION_ONLY=replaced"})
	output := r.Start(
		constant.Shell,
		constant.ShellCommand,
		"echo $INTEGRATION_ONLY:$HOME",
	)
	assert.String(t, "replaced:\n", output)
}

func TestRunInput(t *testing.T) {
	windowsSkip(t)
	r := run.New()
	r.Input = strings.NewReader("first\nsecond\n")
	output := r.Start("cat")
	assert.True(t, r.Error == nil)
	assert.String(t, "first\nsecond\n", output)
	assert.String(t, "first\nsecond\n", r.OutputString)
}

func TestRunInputWithWriters(t *testing.T) {
	windowsSkip(t)
	var stderr bytes.Buffer
	r := run.New()
	r.NoPanic()
	r.Input = strings.NewReader("banana\napple\ncherry\n")
	r.Writers(nil, &stderr)
	output := r.Start("sort")
	assert.True(t, r.Error == nil)
	assert.String(t, "apple\nbanana\ncherry\n", output)
	assert.String(t, "", stderr.String())
}

func TestRunPipe(t *testing.T) {
	windowsSkip(t)
	stdout, stderr := run.New().Pipe("pipe input\n", "cat")
	assert.String(t, "pipe input\n", stdout)
	assert.String(t, "", stderr)
}

func TestRunWritersStdout(t *testing.T) {
	windowsSkip(t)
	var stdout bytes.Buffer
	r := run.New()
	r.Writers(&stdout, nil)
	output := r.Start(
		constant.Shell,
		constant.ShellCommand,
		"echo out; echo err >&2",
	)
	assert.String(t, "out\n", stdout.String())
	assert.String(t, "", output)
	assert.String(t, "", r.OutputString)
	assert.String(t, "err\n", r.ErrorString)
}

func TestRunWritersStderr(t *testing.T) {
	windowsSkip(t)
	var stderr bytes.Buffer
	r := run.New()
	r.Writers(nil, &stderr)
	output := r.Start(
		constant.Shell,
		constant.ShellCommand,
		"echo out; echo err >&2",
	)
	assert.String(t, "out\n", output)
	assert.String(t, "out\n", r.OutputString)
	assert.String(t, "err\n", stderr.String())
	assert.String(t, "", r.ErrorString)
}

func TestProcessKillIsNotAnError(t *testing.T) {
	windowsSkip(t)
	p := run.New().Open("sleep", "30")
	assert.Nil(t, p.Kill())
	assert.Integer(t, -1, p.ExitCode())
}
