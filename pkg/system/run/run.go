package run

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/system/result"
	"io"
)

type Run struct {
	result.Result
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
}
