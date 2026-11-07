package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Prometheus struct {
	option.Web
	Inventory *instance.Inventory
	Client    func(*instance.Instance) (face.PrometheusClient, error)
}
