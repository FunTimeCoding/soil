package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/embed_counter"
	"testing"
)

func newCounter(t *testing.T) *embed_counter.Counter {
	t.Helper()
	result, e := embed_counter.NewEnvironment()
	assert.FatalOnError(t, e)
	t.Cleanup(func() { errors.PanicClose(result) })

	return result
}
