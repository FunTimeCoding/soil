package option

import (
	band "github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

type Band struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
	Engine        func(*instance.Instance) (band.BandClient, error)
	Netbox        func() face.Client
}
