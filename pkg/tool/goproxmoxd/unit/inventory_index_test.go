package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
	"testing"
)

func TestInventoryRejectsSharedIndex(t *testing.T) {
	i := inventory.New(
		inventory.Instance{Name: "first", Index: 0},
		inventory.Instance{Name: "second", Index: 0},
	)
	e := i.Validate()
	assert.Error(t, e)
	assert.StringContains(t, "used by both first and second", e.Error())
}

func TestInventoryAcceptsDistinctIndexes(t *testing.T) {
	i := inventory.New(
		inventory.Instance{Name: "first", Index: 0},
		inventory.Instance{Name: "second", Index: 1},
	)
	assert.Nil(t, i.Validate())
	assert.Integer(t, 1, i.Index("second"))
}

func TestInventorySingleInstanceNeedsNoIndex(t *testing.T) {
	assert.Nil(t, inventory.New(inventory.Instance{Name: "only"}).Validate())
}
