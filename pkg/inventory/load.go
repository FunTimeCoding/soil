package inventory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/inventory/face"
	"github.com/funtimecoding/soil/pkg/system"
	"go.yaml.in/yaml/v3"
)

func Load[T any, P face.Entry[T]](path string) *Inventory[T, P] {
	result := &Inventory[T, P]{}
	errors.PanicOnError(yaml.Unmarshal(system.ReadBytesUnsafe(path), result))
	errors.PanicOnError(result.Validate())

	return result
}
