package inventory

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"

func New(instances ...instance.Instance) *Inventory {
	return &Inventory{Instances: instances}
}
