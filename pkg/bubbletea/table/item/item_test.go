package item

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestItem(t *testing.T) {
	assert.NotNil(t, New(true))
}
