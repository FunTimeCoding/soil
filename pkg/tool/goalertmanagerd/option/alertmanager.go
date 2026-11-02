package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

type Alertmanager struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
	Alertmanager  func(*instance.Instance) (face.AlertmanagerClient, error)
	Prometheus    func(*instance.Instance) (face.PrometheusClient, error)
}
