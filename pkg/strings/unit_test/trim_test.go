package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/trim"
	"testing"
)

func TestNewLine(t *testing.T) {
	assert.String(t, "hello", trim.NewLine("hello\n"))
}
