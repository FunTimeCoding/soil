package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestInventoryItem(t *testing.T) {
	assert.NotNil(t, inventory_item.New(&netbox.InventoryItem{}))
}
