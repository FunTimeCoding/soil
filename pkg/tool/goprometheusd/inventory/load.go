package inventory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"gopkg.in/yaml.v3"
)

func Load(path string) *Inventory {
	result := &Inventory{}
	errors.PanicOnError(yaml.Unmarshal(system.ReadBytesUnsafe(path), result))

	return result
}
