package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/table/item"
	"testing"
)

func TestItem(t *testing.T) {
	assert.NotNil(t, item.New(true))
}
