package option

import (
	band "github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Band struct {
	option.Web
	Inventory *instance.Inventory
	Engine    func(*instance.Instance) (band.BandClient, error)
	Netbox    func() face.Client
}
