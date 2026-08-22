package inventory

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"

type Inventory struct {
	Instances []instance.Instance `yaml:"instances"`
}
