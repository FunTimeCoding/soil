package option

import (
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

type Band struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
	Netbox        func() face.Client
}
