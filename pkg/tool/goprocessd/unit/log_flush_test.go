package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/log"
	"io"
	"os"
	"testing"
	"time"
)

func TestLineWithoutNewlineIsFlushedAfterTimeout(t *testing.T) {
	reader, writer, e := os.Pipe()
	assert.FatalOnError(t, e)
	defer errors.PanicClose(reader)
	l := log.New(writer, "alfa", 0, 4)
	_, f := l.Write([]byte("partial"))
	assert.FatalOnError(t, f)
	time.Sleep(50 * time.Millisecond)
	assert.FatalOnError(t, writer.Close())
	output, g := io.ReadAll(reader)
	assert.FatalOnError(t, g)
	assert.StringContains(t, "partial\n", string(output))
}
