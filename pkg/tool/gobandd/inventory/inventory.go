package inventory

import "github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"

type Inventory struct {
	Instances []instance.Instance `yaml:"instances"`
}
