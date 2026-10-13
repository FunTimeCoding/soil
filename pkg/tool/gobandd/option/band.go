package option

import "github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"

type Band struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
}
