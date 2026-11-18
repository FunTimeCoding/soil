package option

import (
	"github.com/funtimecoding/soil/pkg/nut/face"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Exporter struct {
	Metric    option.Web
	Inventory *instance.Inventory
	Verbose   bool
	Source    func(i *instance.Instance) face.Client
}
