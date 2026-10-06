package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/terminal"
)

type fixture struct {
	terminal *terminal.Terminal
	output   *bytes.Buffer
	failure  *bytes.Buffer
	ends     *commandEnds
	exits    []int
}
