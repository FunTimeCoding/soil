package option

import (
	"github.com/funtimecoding/soil/pkg/postgres/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Postgres struct {
	option.Web
	Inventory *instance.Inventory
	Client    func(*instance.Instance) (face.PostgresClient, error)
}
