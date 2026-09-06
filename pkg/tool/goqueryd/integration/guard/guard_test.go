package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/status")
	c.VerifyOpen(t, constant.DashboardPath)
	c.VerifyModelContext(t)
}
