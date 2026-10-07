package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"testing"
)

func TestInventoryRejectsSharedIndex(t *testing.T) {
	i := inventory.New(
		instance.Instance{Name: "first", Index: 0},
		instance.Instance{Name: "second", Index: 0},
	)
	e := i.Validate()
	assert.Error(t, e)
	assert.StringContains(t, "used by both first and second", e.Error())
}

func TestInventoryAcceptsDistinctIndexes(t *testing.T) {
	i := inventory.New(
		instance.Instance{Name: "first", Index: 0},
		instance.Instance{Name: "second", Index: 1},
	)
	assert.Nil(t, i.Validate())
	assert.Integer(t, 1, i.Index("second"))
}

func TestInventorySingleInstanceNeedsNoIndex(t *testing.T) {
	assert.Nil(t, inventory.New(instance.Instance{Name: "only"}).Validate())
}
