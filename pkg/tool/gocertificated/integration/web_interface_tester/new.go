package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/integration/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{Server: base.New(t)}
}
