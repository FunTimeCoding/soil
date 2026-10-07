package convert

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"

func Instance(
	i *instance.Instance,
	active bool,
) *SlimInstance {
	return &SlimInstance{
		Name:   i.Name,
		Host:   i.Host,
		Port:   i.Port,
		Active: active,
	}
}
