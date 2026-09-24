package assert

import (
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"testing"
)

func Capture(
	t *testing.T,
	f func(),
) string {
	t.Helper()
	original := os.Stdout
	reader, writer, e := os.Pipe()
	FatalOnError(t, e)
	os.Stdout = writer
	done := make(chan string)
	go func() {
		done <- string(system.ReadAll(reader))
	}()
	f()
	os.Stdout = original
	FatalOnError(t, writer.Close())

	return <-done
}
