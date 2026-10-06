package inventory

import "github.com/funtimecoding/soil/pkg/source/inventory/module"

func (i *Inventory) Add(
	name string,
	directory string,
) {
	i.Modules = append(
		i.Modules,
		module.Module{Name: name, Directory: directory},
	)
}
