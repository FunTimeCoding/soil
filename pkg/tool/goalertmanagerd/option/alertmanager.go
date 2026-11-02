package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	prometheusFace "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

type Alertmanager struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
	Alertmanager  func(*instance.Instance) (face.Client, error)
	Prometheus    func(*instance.Instance) (prometheusFace.Client, error)
}
