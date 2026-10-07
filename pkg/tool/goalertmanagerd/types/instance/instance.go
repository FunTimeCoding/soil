package instance

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance/connection"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance/prometheus"
)

type Instance struct {
	Name         string                `yaml:"name"`
	Alertmanager connection.Connection `yaml:"alertmanager"`
	Prometheus   prometheus.Prometheus `yaml:"prometheus"`
}
