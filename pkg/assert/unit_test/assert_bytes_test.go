package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestBytes(t *testing.T) {
	assert.Bytes(t, []byte("123"), []byte("123"))
}
