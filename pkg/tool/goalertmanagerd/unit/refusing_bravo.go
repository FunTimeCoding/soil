package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"testing"
)

func refusingBravo(t *testing.T) *service.Service {
	t.Helper()
	bravo := mock_client.New()
	bravo.Fail(validation.New("bad request"))

	return service_tester.New(
		t,
		map[string]*mock_client.Client{
			"alfa":  mock_client.New(),
			"bravo": bravo,
		},
		map[string]*prometheusMock.Client{
			"alfa":  prometheusMock.New(),
			"bravo": prometheusMock.New(),
		},
	)
}
