package option

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}
