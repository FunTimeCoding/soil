package publish_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/service_tester"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	o := service_tester.New(t)

	return &Tester{Service: o.Service(), Forge: o.Forge, t: t}
}
