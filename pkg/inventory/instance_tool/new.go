package instance_tool

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/face"
	"github.com/funtimecoding/soil/pkg/inventory/selection"
)

func New[T any, P face.Entry[T]](
	i *inventory.Inventory[T, P],
	s *selection.Selection,
) *Tool[T, P] {
	return &Tool[T, P]{inventory: i, selection: s}
}
