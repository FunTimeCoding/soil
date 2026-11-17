package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"testing"
)

func NewRefusingBravo(t *testing.T) *service.Service {
	t.Helper()
	o := New(t, WithInstances("alfa", "bravo"))
	o.Bravo().Alertmanager.Fail(validation.New("bad request"))

	return o.Service()
}
