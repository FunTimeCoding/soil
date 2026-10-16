package connection

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance/endpoint"

type Connection struct {
	endpoint.Endpoint `yaml:",inline"`
	User              string            `yaml:"user"`
	Password          string            `yaml:"password"`
	Link              endpoint.Endpoint `yaml:"link"`
}
