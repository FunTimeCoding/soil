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
	reporter       face.Reporter
	reporterLabel  string
	registry       face.ProcessRegistry
	Directory      string
	Panic          bool
	Verbose        bool
	OutputString   string
	ErrorString    string
	Error          error
	Exit           int
}
