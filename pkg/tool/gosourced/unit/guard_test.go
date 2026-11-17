package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	c := base.New(
		t,
		base.WithFixture(service_tester.ServiceTestdata("cross-package/src")),
	).Served()
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/diff")
	c.VerifyGuarded(t, "/api/usage")
	c.VerifyModelContext(t)
}
