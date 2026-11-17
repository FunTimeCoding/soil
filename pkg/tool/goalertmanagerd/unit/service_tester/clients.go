package service_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
)

type Clients struct {
	Alertmanager *mock_client.Client
	Prometheus   *prometheusMock.Client
}
