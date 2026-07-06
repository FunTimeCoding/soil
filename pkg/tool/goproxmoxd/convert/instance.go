package convert

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"

func Instance(
	i *inventory.Instance,
	active bool,
) *SlimInstance {
	return &SlimInstance{
		Name:   i.Name,
		Host:   i.Host,
		Active: active,
	}
}
