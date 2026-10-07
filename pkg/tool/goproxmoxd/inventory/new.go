package inventory

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

func New(instances ...instance.Instance) *Inventory {
	return &Inventory{Instances: instances}
}
