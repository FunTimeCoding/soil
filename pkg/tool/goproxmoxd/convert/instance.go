package convert

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

func Instance(
	i *instance.Instance,
	active bool,
) *SlimInstance {
	return &SlimInstance{Name: i.Name, Host: i.Host, Active: active}
}
