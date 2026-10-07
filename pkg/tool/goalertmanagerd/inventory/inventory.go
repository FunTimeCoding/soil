package inventory

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"

type Inventory struct {
	Instances []instance.Instance `yaml:"instances"`
}
