package notation

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMarshallIndent(t *testing.T) {
	assert.Any(t, `"a"`, MarshalIndent("a"))
}
