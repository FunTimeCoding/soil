package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"testing"
)

func TestMarshallIndent(t *testing.T) {
	assert.Any(t, `"a"`, notation.MarshalIndent("a"))
}
