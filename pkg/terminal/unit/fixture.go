package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/unit/command_end_recorder"
)

type Fixture struct {
	terminal *terminal.Terminal
	output   *bytes.Buffer
	failure  *bytes.Buffer
	ends     *command_end_recorder.Recorder
	exits    []int
}
