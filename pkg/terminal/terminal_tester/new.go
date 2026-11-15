package terminal_tester

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/command_end_recorder"
)

func New() *Tester {
	result := &Tester{
		Output:  &bytes.Buffer{},
		Failure: &bytes.Buffer{},
		Ends:    command_end_recorder.New(),
	}
	result.Terminal = terminal.NewWith(
		result.Ends,
		result.Output,
		result.Failure,
		func(code int) { result.Exits = append(result.Exits, code) },
	)

	return result
}
