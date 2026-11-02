package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
)

type Prometheus struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
	Client        func(*instance.Instance) (face.Client, error)
}
