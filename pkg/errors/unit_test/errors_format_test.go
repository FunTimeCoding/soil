package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.String(t, `a: "b"`, errors.Format("a", "b").Error())
}
