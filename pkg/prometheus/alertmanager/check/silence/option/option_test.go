package option

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestOption(t *testing.T) {
	assert.NotNil(t, New())
}
