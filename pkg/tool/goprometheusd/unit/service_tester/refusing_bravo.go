package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"testing"
)

func RefusingBravo(t *testing.T) *service.Service {
	t.Helper()
	bravo := mock_client.New()
	bravo.Failure = validation.New("bad query")

	return New(
		t,
		map[string]*mock_client.Client{
			"alfa":  mock_client.New(),
			"bravo": bravo,
		},
	)
}
