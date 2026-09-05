package run

import (
	"bytes"
	"errors"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/command"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os/exec"
)

func (r *Run) Start(s ...string) string {
	if r.Verbose {
		fmt.Printf("Run: %s\n", join.Space(s...))
	}

	c := r.build(s...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	if r.stdout == nil {
		c.Stdout = &stdout
	}

	if r.stderr == nil {
		c.Stderr = &stderr
	}

	e := r.startAndWait(c)

	if r.stdout == nil {
		r.OutputString = stdout.String()
	}

	if r.stderr == nil {
		r.ErrorString = stderr.String()
	}

	if e != nil {
		r.Error = command.New(
			join.Space(s...),
			r.OutputString,
			r.ErrorString,
			e,
		)
	}

	if r.Verbose {
		r.Print()
	}

	if f, okay := errors.AsType[*exec.ExitError](e); okay {
		r.Exit = f.ExitCode()
	}

	if r.Panic && e != nil {
		library.PanicOnError(r.Error)
	}

	return r.OutputString
}
