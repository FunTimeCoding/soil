package instance_tool

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/face"
)

func New[T any, P face.Entry[T]](
	i *inventory.Inventory[T, P],
	s *selection.Selection,
	r soilFace.Reporter,
) *Tool[T, P] {
	return &Tool[T, P]{inventory: i, selection: s, reporter: r}
}
