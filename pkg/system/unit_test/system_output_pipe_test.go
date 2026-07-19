package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestOutputStream(t *testing.T) {
	s := run.New().Stream("echo", "hello")
	assert.String(t, "hello\n", string(system.ReadAll(s.Reader())))
	errors.PanicOnError(s.Wait())
}
