package option

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"

type Prometheus struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
}
