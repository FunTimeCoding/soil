package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	alertmanagerMock "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"testing"
)

func NewRefusingBravo(t *testing.T) *service.Service {
	t.Helper()
	bravo := alertmanagerMock.New()
	bravo.Fail(validation.New("bad request"))

	return New(
		t,
		map[string]*alertmanagerMock.Client{
			"alfa":  alertmanagerMock.New(),
			"bravo": bravo,
		},
		map[string]*prometheusMock.Client{
			"alfa":  prometheusMock.New(),
			"bravo": prometheusMock.New(),
		},
	)
}
