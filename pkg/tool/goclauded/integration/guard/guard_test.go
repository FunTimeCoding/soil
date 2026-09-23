package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	c := s.Server
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/sessions")
	c.VerifyGuarded(t, "/api/status")
	c.VerifyGuarded(t, "/api/channel")
	c.VerifyGuarded(t, "/api/channel/callsign")
	c.VerifyOpen(t, constant.DashboardPath)
	c.VerifyOpen(t, constant.StatusPath)
	c.VerifyOpen(t, "/event")
	c.VerifyModelContext(t)
}
