package convert

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/inventory"

func Instance(
	i *inventory.Instance,
	active bool,
) *SlimInstance {
	return &SlimInstance{
		Name:   i.Name,
		Host:   i.Host,
		Port:   i.Port,
		Active: active,
	}
}
