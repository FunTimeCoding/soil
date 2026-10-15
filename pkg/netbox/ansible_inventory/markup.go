package ansible_inventory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"go.yaml.in/yaml/v3"
)

func (i *Inventory) Markup() string {
	children := map[string]any{}

	for group, hosts := range i.groups {
		children[group] = map[string]any{constant.InventoryHosts: hosts}
	}

	m, e := yaml.Marshal(
		map[string]any{
			constant.InventoryAll: map[string]any{
				constant.InventoryChildren: children,
			},
		},
	)
	errors.PanicOnError(e)

	return string(m)
}
