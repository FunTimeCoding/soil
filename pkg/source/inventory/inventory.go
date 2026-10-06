package inventory

import "github.com/funtimecoding/soil/pkg/source/inventory/module"

type Inventory struct {
	Modules []module.Module `yaml:"modules"`
}
