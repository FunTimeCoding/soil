package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"testing"
)

func TestInventoryRejectsSharedIndex(t *testing.T) {
	e := instance.ValidateIndexes(
		inventory.New[instance.Instance](
			instance.Instance{Name: "first", Index: 0},
			instance.Instance{Name: "second", Index: 0},
		),
	)
	assert.Error(t, e)
	assert.StringContains(t, "used by both first and second", e.Error())
}

func TestInventoryAcceptsDistinctIndexes(t *testing.T) {
	i := inventory.New[instance.Instance](
		instance.Instance{Name: "first", Index: 0},
		instance.Instance{Name: "second", Index: 1},
	)
	assert.Nil(t, instance.ValidateIndexes(i))
	assert.Integer(t, 1, instance.Index(i, "second"))
}

func TestInventorySingleInstanceNeedsNoIndex(t *testing.T) {
	assert.Nil(
		t,
		instance.ValidateIndexes(
			inventory.New[instance.Instance](instance.Instance{Name: "only"}),
		),
	)
}
