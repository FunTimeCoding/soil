package example_list

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestGroceryList(t *testing.T) {
	assert.NotNil(t, New())
}
