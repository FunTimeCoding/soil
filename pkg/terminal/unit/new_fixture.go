package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/unit/command_end_recorder"
)

func newFixture() *Fixture {
	result := &Fixture{
		output:  &bytes.Buffer{},
		failure: &bytes.Buffer{},
		ends:    command_end_recorder.New(),
	}
	result.terminal = terminal.NewWith(
		result.ends,
		result.output,
		result.failure,
		func(code int) { result.exits = append(result.exits, code) },
	)

	return result
}
