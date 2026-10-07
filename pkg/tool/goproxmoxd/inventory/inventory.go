package inventory

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

type Inventory struct {
	Instances []instance.Instance `yaml:"instances"`
}
