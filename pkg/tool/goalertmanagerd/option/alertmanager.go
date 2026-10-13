package option

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"

type Alertmanager struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
}
