package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/example_list"
	"testing"
)

func TestGroceryList(t *testing.T) {
	assert.NotNil(t, example_list.New())
}
