package ansible_inventory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"go.yaml.in/yaml/v3"
)

func (i *Inventory) Markup() string {
	children := map[string]any{}

	for group, hosts := range i.groups {
		entries := map[string]any{}

		for host, address := range hosts {
			variables := map[string]any{}

			if address != "" {
				variables[constant.InventoryAddress] = address
			}

			entries[host] = variables
		}

		children[group] = map[string]any{constant.InventoryHosts: entries}
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
