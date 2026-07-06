package expression

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestExpression(t *testing.T) {
	e := New([]string{upper.Alfa}, []string{upper.Delta})
	assert.True(t, e.Check([]string{upper.Alfa}))
	assert.True(t, e.Check([]string{upper.Alfa, upper.Bravo}))
	assert.False(t, e.Check([]string{upper.Alfa, upper.Delta}))
}
