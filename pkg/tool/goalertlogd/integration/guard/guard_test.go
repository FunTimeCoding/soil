package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/alerts")
	c.VerifyGuarded(t, "/api/status")
	c.VerifyOpen(t, constant.DashboardPath)
	c.VerifyOpen(t, "/alerts")
	c.VerifyModelContext(t)
}
