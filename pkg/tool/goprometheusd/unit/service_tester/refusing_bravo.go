package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"testing"
)

func RefusingBravo(t *testing.T) *service.Service {
	t.Helper()
	o := New(t, WithInstances("alfa", "bravo"))
	o.Bravo().Failure = validation.New("bad query")

	return o.Service()
}
