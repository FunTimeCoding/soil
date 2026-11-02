package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/service_tester"
	"testing"
)

func refusingBravo(t *testing.T) *service.Service {
	t.Helper()
	bravo := mock_client.New()
	bravo.Failure = validation.New("bad query")

	return service_tester.New(
		t,
		map[string]*mock_client.Client{
			"alfa":  mock_client.New(),
			"bravo": bravo,
		},
	)
}
