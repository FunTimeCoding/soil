package testutil

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"os"
)

func CaptureStdout(f func()) string {
	original := os.Stdout
	reader, writer, e := os.Pipe()
	errors.PanicOnError(e)
	os.Stdout = writer
	f()
	errors.PanicClose(writer)
	os.Stdout = original
	captured, e := io.ReadAll(reader)
	errors.PanicOnError(e)

	return string(captured)
}
