package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	c := base.New(t).Served()
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/events")
	c.VerifyGuarded(t, "/api/status")
}
