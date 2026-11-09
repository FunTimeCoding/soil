package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"testing"
	"time"
)

func NewWith(
	t *testing.T,
	clock func() time.Time,
) *Tester {
	t.Helper()

	return &Tester{Server: base.NewWith(t, clock), t: t}
}
