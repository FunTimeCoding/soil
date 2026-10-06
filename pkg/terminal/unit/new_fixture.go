package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/terminal"
)

func newFixture() *fixture {
	result := &fixture{
		output:  &bytes.Buffer{},
		failure: &bytes.Buffer{},
		ends:    &commandEnds{},
	}
	result.terminal = terminal.NewWith(
		result.ends,
		result.output,
		result.failure,
		func(code int) { result.exits = append(result.exits, code) },
	)

	return result
}
