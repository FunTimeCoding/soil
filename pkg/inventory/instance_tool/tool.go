package instance_tool

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/face"
	"github.com/funtimecoding/soil/pkg/inventory/selection"
)

type Tool[T any, P face.Entry[T]] struct {
	inventory *inventory.Inventory[T, P]
	selection *selection.Selection
}
