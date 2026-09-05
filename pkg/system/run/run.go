package run

import (
	"github.com/funtimecoding/soil/pkg/face"
	"io"
)

type Run struct {
	environment    []string
	replaceEnviron bool
	processGroup   bool
	stdio          bool
	stdout         io.Writer
	stderr         io.Writer
	registry       face.ProcessRegistry
	Directory      string
	Input          io.Reader
	Panic          bool
	Verbose        bool
	OutputString   string
	ErrorString    string
	Error          error
	Exit           int
}
