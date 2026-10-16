package instance

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance/connection"

type Instance struct {
	Name         string                `yaml:"name"`
	Alertmanager connection.Connection `yaml:"alertmanager"`
	Prometheus   connection.Connection `yaml:"prometheus"`
}
