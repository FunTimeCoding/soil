package inventory

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"

type Inventory struct {
	Instances []instance.Instance `yaml:"instances"`
}
