package terminal_tester

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/command_end_recorder"
)

type Tester struct {
	Terminal *terminal.Terminal
	Output   *bytes.Buffer
	Failure  *bytes.Buffer
	Ends     *command_end_recorder.Recorder
	Exits    []int
}
