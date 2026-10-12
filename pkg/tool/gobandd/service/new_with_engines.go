package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory"
	"maps"
)

func NewWithEngines(
	i *inventory.Inventory,
	engines map[string]face.Engine,
) *Service {
	result := New(i)
	maps.Copy(result.engines, engines)

	return result
}
