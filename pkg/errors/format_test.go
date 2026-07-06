package errors

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.String(t, `a: "b"`, Format("a", "b").Error())
}
