package option

import "github.com/funtimecoding/soil/pkg/tool/gobandd/inventory"

type Band struct {
	Address       string
	ServiceTokens []string
	Inventory     *inventory.Inventory
}
